package cmd

import (
	"path/filepath"
	"fmt"
	"encoding/json"
	"io/fs"
	"strings"
	"os"
	"scai-gen/fileio"

	"github.com/adityasaky/in-toto-attestation-verifier/verifier"
	"github.com/secure-systems-lab/go-securesystemslib/dsse"
	//ita "github.com/in-toto/attestation/go/v1"
	//scai "github.com/in-toto/attestation/go/predicates/scai/v0"
	"github.com/spf13/cobra"
	//"google.golang.org/protobuf/encoding/protojson"
	//"google.golang.org/protobuf/types/known/structpb"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify one or more JSON-encoded SCAI Attribute Reports according to an in-toto ITE-10 layout",
}

var attrCmd = &cobra.Command{
	Use:   "attr",
	Args:  cobra.MinimumNArgs(1),
	Short: "Verify the attributes of one or more JSON-encoded SCAI Attribute Reports according to an in-toto ITE-10 layout (uses in-toto-attestation-verifier)",
	RunE:  verifyLayout,
}

var evCmd = &cobra.Command{
	Use:   "ev",
	Args:  cobra.MinimumNArgs(1),
	Short: "Verify the evidence for one or more JSON-encoded SCAI Attribute Reports according to an in-toto ITE-10 layout",
	RunE:  verifyEvidence,
}

var (
	layoutFile  string
	evidenceDir string
)

func init() {
	verifyCmd.AddCommand(attrCmd)
	verifyCmd.AddCommand(evCmd)
}

func init() {
	attrCmd.Flags().StringVarP(
		&layoutFile,
		"layout",
		"l",
		"",
		"The filename of the YAML-encoded in-toto Layout",
	)
	attrCmd.MarkFlagRequired("layout")
}

func init() {
	evCmd.Flags().StringVarP(
		&evidenceDir,
		"evidence-dir",
		"e",
		"",
		"The directory containing evidence files",
	)
	evCmd.MarkFlagRequired("evidence-dir")
}

func verifyLayout(cmd *cobra.Command, args []string) error {

	layout, err := verifier.LoadLayout(layoutFile)
	if err != nil {
		return err
	}

	attestations := map[string]*dsse.Envelope{}
	for _, attestationPath := range args {
		name := filepath.Base(attestationPath)
		envBytes, err := os.ReadFile(attestationPath)
		if err != nil {
			return err
		}

		envelope := &dsse.Envelope{}
		if err := json.Unmarshal(envBytes, envelope); err != nil {
			return err
		}

		attestationName := strings.TrimSuffix(name, ".json")
		attestations[attestationName] = envelope
		fmt.Println("Found attestation ", attestationName)
	}

	parameters := map[string]string{}

	fmt.Println("Verifying ", len(attestations), "attestation(s)")

	return verifier.Verify(layout, attestations, parameters)
}

func verifyEvidence(cmd *cobra.Command, args []string) error {

	evidence := map[string][]byte{}
	err := filepath.Walk(evidenceDir, func(path string, info fs.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			return fileio.ReadFileIntoMap(path, evidence)
		}
		return err
	})
	if err != nil {
		return err
	}

	for file, _ := range evidence {
		fmt.Println("Found evidence file ", file)
	}

	/* 
	statements := map[string]*ita.Statement{}
	for _, attestationPath := range args {
		name := filepath.Base(attestationPath)
		stBytes, err := os.ReadFile(attestationPath)
		if err != nil {
			return err
		}

		envelope := &ita.Statement{}
		if err := json.Unmarshal(envBytes, envelope); err != nil {
			return err
		}

		attestationName := strings.TrimSuffix(name, ".json")
		attestations[attestationName] = envelope
		fmt.Println("Found attestation ", attestationName)
	}

	parameters := map[string]string{}

	fmt.Println("Verifying ", len(attestations), "attestation(s)")
	*/

	return nil
}
