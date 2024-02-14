package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/in-toto/scai-demos/scai-gen/pkg/fileio"
	"github.com/in-toto/scai-demos/scai-gen/pkg/policy"

	"github.com/in-toto/attestation-verifier/verifier"
	scai "github.com/in-toto/attestation/go/predicates/scai/v0"
	ita "github.com/in-toto/attestation/go/v1"
	"github.com/secure-systems-lab/go-securesystemslib/dsse"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/yaml.v3"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check one or more JSON-encoded SCAI attestations",
}

var layoutCmd = &cobra.Command{
	Use:   "layout",
	Args:  cobra.MinimumNArgs(1),
	Short: "Check the attributes of one or more JSON-encoded SCAI attestations according to an in-toto ITE-10 layout",
	RunE:  checkLayout,
}

var evCmd = &cobra.Command{
	Use:   "evidence",
	Args:  cobra.ExactArgs(1),
	Short: "Check the evidence for a JSON-encoded SCAI attestation according to an evidence policy",
	RunE:  checkEvidence,
}

var (
	layoutFile  string
	evidenceDir string
	policyFile  string
)

func init() {
	checkCmd.AddCommand(layoutCmd)
	checkCmd.AddCommand(evCmd)
}

func init() {
	layoutCmd.Flags().StringVarP(
		&layoutFile,
		"layout",
		"l",
		"",
		"The filename of the YAML-encoded in-toto Layout",
	)
	layoutCmd.MarkFlagRequired("layout") //nolint:errcheck
}

func init() {
	evCmd.Flags().StringVarP(
		&evidenceDir,
		"evidence-dir",
		"e",
		"",
		"The directory containing evidence files",
	)
	evCmd.MarkFlagRequired("evidence-dir") //nolint:errcheck

	evCmd.Flags().StringVarP(
		&policyFile,
		"policy-file",
		"p",
		"",
		"The filename of the policy file",
	)
	evCmd.MarkFlagRequired("policy-file") //nolint:errcheck
}

func checkLayout(_ *cobra.Command, args []string) error {
	layout, err := verifier.LoadLayout(layoutFile)
	if err != nil {
		return err
	}

	attestations := map[string]*dsse.Envelope{}
	for _, attestationPath := range args {
		name := filepath.Base(attestationPath)

		envelope, err := fileio.ReadDSSEFile(attestationPath)
		if err != nil {
			return err
		}

		attestationName := strings.TrimSuffix(name, ".json")
		attestations[attestationName] = envelope
		fmt.Println("Found attestation ", attestationName)
	}

	parameters := map[string]string{}

	fmt.Println("Checking ", len(attestations), "attestation(s)")

	return verifier.Verify(layout, attestations, parameters)
}

func checkEvidence(_ *cobra.Command, args []string) error {
	attestationPath := args[0]
	fmt.Println("Reading attestation file", attestationPath)

	envBytes, err := os.ReadFile(attestationPath)
	if err != nil {
		return err
	}

	fmt.Println("Reading policy file", policyFile)

	policyBytes, err := os.ReadFile(policyFile)
	if err != nil {
		return err
	}

	evPolicy := &policy.SCAIEvidencePolicy{}
	if err := yaml.Unmarshal(policyBytes, evPolicy); err != nil {
		return err
	}

	fmt.Println("Checking attestation matches ID in policy")

	if !policy.MatchDigest(evPolicy.AttestationID, envBytes) {
		return fmt.Errorf("attestation does not match attestation ID in policy")
	}

	// now, let's get the Statement
	// we don't call fileio.ReadDSSEFile to ensure we are evaluating
	// over the matched attestation
	envelope := &dsse.Envelope{}
	if err := json.Unmarshal(envBytes, envelope); err != nil {
		return err
	}

	statement, err := getStatementDSSEPayload(envelope)
	if err != nil {
		return err
	}

	fmt.Println("Collecting all evidence files")

	evidenceFiles, err := getAllEvidenceFiles(evidenceDir)
	if err != nil {
		return fmt.Errorf("failed read evidence files in directory %s: %w", evidenceDir, err)
	}

	if statement.GetPredicateType() != "https://in-toto.io/attestation/scai/attribute-report/v0.2" {
		return fmt.Errorf("evidence checking only supported for SCAI attestations")
	}

	report, err := pbStructToSCAI(statement.GetPredicate())
	if err != nil {
		return err
	}

	// validate the report
	if err := report.Validate(); err != nil {
		return fmt.Errorf("malformed SCAI Attribute Report: %w", err)
	}

	// order attribute assertions by evidence name
	// FIXME: for now assume that there's a 1:1 mapping of AA to evidence
	attrAssertions := map[string]*scai.AttributeAssertion{}
	for _, attrAssertion := range report.GetAttributes() {
		if ev := attrAssertion.GetEvidence(); ev != nil {
			attrAssertions[ev.GetName()] = attrAssertion
		}
	}

	fmt.Println("Checking policy rules...")

	for _, check := range evPolicy.Inspections {
		rules := check.ExpectedAttributes
		if len(rules) == 0 {
			return fmt.Errorf("no rules for check %s", check.Name)
		}

		attrAssertion, ok := attrAssertions[check.Name]
		if !ok {
			return fmt.Errorf("attestation evidence missing %s", check.Name)
		}

		fmt.Println("Validating attribute assertion format")
		if err := attrAssertion.Validate(); err != nil {
			return fmt.Errorf("malformed attribute assertion in attestation: %w", err)
		}

		ev := attrAssertion.GetEvidence()

		evContent, ok := evidenceFiles[ev.GetName()]
		if !ok {
			return fmt.Errorf("evidence file to check not found")
		}

		fmt.Println("Checking evidence content according to policy rules...")

		switch ev.GetMediaType() {
		case "text/plain":
			err := policy.ApplyPlaintextRules(string(evContent), attrAssertion, rules)
			if err != nil {
				return fmt.Errorf("plaintext policy check failed: %w", err)
			}

		case "application/vnd.in-toto+dsse":
			evEnv := &dsse.Envelope{}
			if err := json.Unmarshal(evContent, evEnv); err != nil {
				return err
			}

			evStatement, err := getStatementDSSEPayload(evEnv)
			if err != nil {
				return err
			}

			err = policy.ApplyAttestationRules(evStatement, attrAssertion, rules)
			if err != nil {
				return fmt.Errorf("attestation policy check failed: %w", err)
			}

		default:
			return fmt.Errorf("evidence type not supported: %s", ev.GetMediaType())
		}
	}

	fmt.Println("Evidence policy checks successful!")

	return nil
}

func pbStructToSCAI(s *structpb.Struct) (*scai.AttributeReport, error) {
	structJSON, err := protojson.Marshal(s)
	if err != nil {
		return nil, err
	}

	report := &scai.AttributeReport{}
	err = protojson.Unmarshal(structJSON, report)
	if err != nil {
		return nil, err
	}

	return report, nil
}

func getStatementDSSEPayload(envelope *dsse.Envelope) (*ita.Statement, error) {
	stBytes, err := envelope.DecodeB64Payload()
	if err != nil {
		return nil, fmt.Errorf("failed to decode DSSE payload: %w", err)
	}

	statement := &ita.Statement{}
	if err = protojson.Unmarshal(stBytes, statement); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Statement: %w", err)
	}

	return statement, nil
}

func getAllEvidenceFiles(evidenceDir string) (map[string][]byte, error) {
	evidenceMap := map[string][]byte{}
	err := filepath.Walk(evidenceDir, func(path string, info fs.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			return fileio.ReadFileIntoMap(path, evidenceMap)
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return evidenceMap, nil
}
