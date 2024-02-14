package cmd

import (
	"fmt"

	"github.com/in-toto/scai-demos/scai-gen/pkg/fileio"
	"github.com/in-toto/scai-demos/scai-gen/pkg/generators"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/structpb"
)

var rdCmd = &cobra.Command{
	Use:   "rd",
	Short: "Generate a JSON-encoded in-toto Resource Descriptor",
}

var rdFileCmd = &cobra.Command{
	Use:   "file",
	Args:  cobra.ExactArgs(1),
	Short: "Generate a JSON-encoded in-toto Resource Descriptor for a file",
	RunE:  genRdFromFile,
}

var rdRemoteCmd = &cobra.Command{
	Use:   "remote",
	Args:  cobra.ExactArgs(1),
	Short: "Generate a JSON-encoded in-toto Resource Descriptor for a remote resource, identified via a URI, incl. git repos, container images, packages, and services",
	RunE:  genRdForRemote,
}

var (
	digest           string
	downloadLocation string
	hashAlg          string
	mediaType        string
	name             string
	uri              string
	withContent      bool
	annotationsFile  string
)

func init() {
	rdCmd.PersistentFlags().StringVarP(
		&outFile,
		"out-file",
		"o",
		"",
		"Filename to write out the JSON-encoded object",
	)
	rdCmd.MarkPersistentFlagRequired("out-file") //nolint:errcheck

	rdCmd.AddCommand(rdFileCmd)
	rdCmd.AddCommand(rdRemoteCmd)
}

func init() {
	rdFileCmd.Flags().StringVarP(
		&name,
		"name",
		"n",
		"",
		"A name for the local file",
	)

	rdFileCmd.Flags().StringVarP(
		&uri,
		"uri",
		"u",
		"",
		"The URI of the resource",
	)

	rdFileCmd.Flags().StringVarP(
		&downloadLocation,
		"download-location",
		"l",
		"",
		"The download location of the resource (if different from the URI)",
	)

	rdFileCmd.Flags().StringVarP(
		&mediaType,
		"media-type",
		"t",
		"",
		"The media type of the resource",
	)

	rdFileCmd.Flags().StringVarP(
		&annotationsFile,
		"annotations",
		"a",
		"",
		"Filename of a JSON-encoded file containingt annotations for the file",
	)
}

func init() {
	rdRemoteCmd.Flags().StringVarP(
		&name,
		"name",
		"n",
		"",
		"A name for the remote resource or service",
	)

	rdRemoteCmd.Flags().StringVarP(
		&downloadLocation,
		"download-location",
		"l",
		"",
		"The download location of the remote resource (if different from the URI)",
	)

	rdRemoteCmd.Flags().StringVarP(
		&digest,
		"digest",
		"d",
		"",
		"The digest associated with the remote resource (hex-encoded)",
	)

	rdRemoteCmd.Flags().StringVarP(
		&hashAlg,
		"hash-alg",
		"g",
		"",
		"The hash algorithm used to compute the digest associated with the remote resource",
	)
	rdRemoteCmd.MarkFlagsRequiredTogether("digest", "hash-alg")

	rdRemoteCmd.Flags().StringVarP(
		&annotationsFile,
		"annotations",
		"a",
		"",
		"Filename of a JSON-encoded file containingt annotations for the resource",
	)
}

func readAnnotations(filename string) (*structpb.Struct, error) {
	var annotations *structpb.Struct
	if len(filename) > 0 {
		annotations = &structpb.Struct{}
		err := fileio.ReadPbFromFile(filename, annotations)
		if err != nil {
			return nil, err
		}
	}

	return annotations, nil
}

func genRdFromFile(_ *cobra.Command, args []string) error {
	// want to make sure the ResourceDescriptor is a JSON file
	if !fileio.HasJSONExt(outFile) {
		return fmt.Errorf("expected a .json extension for the generated ResourceDescriptor file %s", outFile)
	}

	filename := args[0]

	annotations, err := readAnnotations(annotationsFile)
	if err != nil {
		return fmt.Errorf("error reading annotations file: %w", err)
	}

	rd, err := generators.NewRdForFile(filename, name, uri, hashAlg, withContent, mediaType, downloadLocation, annotations)
	if err != nil {
		return fmt.Errorf("error generating RD: %w", err)
	}

	return fileio.WritePbToFile(rd, outFile, false)
}

func genRdForRemote(_ *cobra.Command, args []string) error {
	// want to make sure the ResourceDescriptor is a JSON file
	if !fileio.HasJSONExt(outFile) {
		return fmt.Errorf("expected a .json extension for the generated ResourceDescriptor file %s", outFile)
	}

	remoteURI := args[0]

	annotations, err := readAnnotations(annotationsFile)
	if err != nil {
		return fmt.Errorf("error reading annotations file: %w", err)
	}

	rd, err := generators.NewRdForRemote(remoteURI, name, hashAlg, digest, downloadLocation, annotations)
	if err != nil {
		return fmt.Errorf("error generating RD: %w", err)
	}

	return fileio.WritePbToFile(rd, outFile, false)
}
