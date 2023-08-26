package cmd

import (
	"fmt"
	"scai-gen/fileio"

	ita "github.com/in-toto/attestation/go/v1"
	scai "github.com/in-toto/attestation/go/predicates/scai/v0"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Args:  cobra.MinimumNArgs(1),
	Short: "Generate a JSON-encoded SCAI Attribute Report",
	RunE:  genAttrReport,
}

var (
	subjectFile  string
	producerFile string
)

func init() {
	reportCmd.Flags().StringVarP(
		&outFile,
		"out-file",
		"o",
		"",
		"Filename to write out the JSON-encoded object",
	)
	reportCmd.MarkFlagRequired("out-file")
	
	reportCmd.Flags().StringVarP(
		&subjectFile,
		"subject",
		"s",
		"",
		"The filename of the JSON-encoded subject resource descriptor",
	)
	reportCmd.MarkFlagRequired("subject")

	reportCmd.Flags().StringVarP(
		&producerFile,
		"producer",
		"p",
		"",
		"The filename of the JSON-encoded producer resource descriptor",
	)
}

func genAttrReport(cmd *cobra.Command, args []string) error {

	var attrAsserts []*scai.AttributeAssertion
	for _, attrAssertPath := range args {
		aa := &scai.AttributeAssertion{}
		err := fileio.ReadPbFromFile(attrAssertPath, aa)
		if err != nil {
			return err
		}

		attrAsserts = append(attrAsserts, aa)
	}

	var producer *ita.ResourceDescriptor
	if len(producerFile) > 0 {
		producer = &ita.ResourceDescriptor{}
		err := fileio.ReadPbFromFile(producerFile, producer)
		if err != nil {
			return err
		}
	}

	// first, generate the SCAI Report
	ar := &scai.AttributeReport{
		Attributes: attrAsserts,
		Producer: producer,
	}

	err := ar.Validate()
	if err != nil {
		return fmt.Errorf("Invalid SCAI attribute report: %w", err)
	}

	// then, plug the Report into an in-toto Statement

	var subject *ita.ResourceDescriptor
	if len(subjectFile) > 0 {
		subject = &ita.ResourceDescriptor{}
		err := fileio.ReadPbFromFile(subjectFile, subject)
		if err != nil {
			return err
		}
	}

	reportJson, err := protojson.Marshal(ar)
	if err != nil {
		return err
	}
	reportStruct := &structpb.Struct{}
	err = protojson.Unmarshal(reportJson, reportStruct)
	if err != nil {
		return err
	}

	statement := &ita.Statement{
		Type:          ita.StatementTypeUri,
		Subject:       []*ita.ResourceDescriptor{subject},
		PredicateType: "https://in-toto.io/attestation/scai/attribute-report/v0.2",
		Predicate:     reportStruct,
	}

	err = statement.Validate()
	if err != nil {
		return fmt.Errorf("Invalid in-toto Statement: %w", err)
	}
	
	return fileio.WritePbToFile(statement, outFile)
}
