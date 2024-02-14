package cmd

import (
	"fmt"

	"github.com/in-toto/scai-demos/scai-gen/pkg/fileio"
	"github.com/in-toto/scai-demos/scai-gen/pkg/generators"

	ita "github.com/in-toto/attestation/go/v1"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/structpb"
)

var assertCmd = &cobra.Command{
	Use:   "assert",
	Args:  cobra.ExactArgs(1),
	Short: "Generate a JSON-encoded SCAI Attribute Assertion",
	RunE:  genAttrAssertion,
}

var (
	targetFile     string
	conditionsFile string
	evidenceFile   string
)

func init() {
	assertCmd.Flags().StringVarP(
		&outFile,
		"out-file",
		"o",
		"",
		"Filename to write out the JSON-encoded object",
	)
	assertCmd.MarkFlagRequired("out-file") //nolint:errcheck

	assertCmd.Flags().StringVarP(
		&targetFile,
		"target",
		"t",
		"",
		"The filename of the JSON-encoded target resource descriptor",
	)

	assertCmd.Flags().StringVarP(
		&conditionsFile,
		"conditions",
		"c",
		"",
		"The filename of the JSON-encoded conditions object",
	)

	assertCmd.Flags().StringVarP(
		&evidenceFile,
		"evidence",
		"e",
		"",
		"The filename of the JSON-encoded evidence resource descriptor",
	)
}

func genAttrAssertion(_ *cobra.Command, args []string) error {
	// want to make sure the AttributeAssertion is a JSON file
	if !fileio.HasJSONExt(outFile) {
		return fmt.Errorf("expected a .json extension for the generated SCAI AttributeAssertion file %s", outFile)
	}

	attribute := args[0]

	var target *ita.ResourceDescriptor
	if len(targetFile) > 0 {
		target = &ita.ResourceDescriptor{}
		err := fileio.ReadPbFromFile(targetFile, target)
		if err != nil {
			return err
		}
	}

	var conditions *structpb.Struct
	if len(conditionsFile) > 0 {
		conditions = &structpb.Struct{}
		err := fileio.ReadPbFromFile(conditionsFile, conditions)
		if err != nil {
			return err
		}
	}

	var evidence *ita.ResourceDescriptor
	if len(evidenceFile) > 0 {
		evidence = &ita.ResourceDescriptor{}
		err := fileio.ReadPbFromFile(evidenceFile, evidence)
		if err != nil {
			return err
		}
	}

	aa, err := generators.NewSCAIAssertion(attribute, target, conditions, evidence)
	if err != nil {
		return fmt.Errorf("error generating SCAI attribute assertion: %w", err)
	}

	return fileio.WritePbToFile(aa, outFile, false)
}
