package cmd

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"github.com/in-toto/scai-demos/scai-gen/fileio"
	"github.com/in-toto/scai-demos/scai-gen/policy"

	ita "github.com/in-toto/attestation/go/v1"
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
	rdCmd.MarkPersistentFlagRequired("out-file")
	
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
	
	rdFileCmd.Flags().BoolVarP(
		&withContent,
		"content",
		"c",
		false,
		"Flag to include the content of the file",
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

func genRdFromFile(cmd *cobra.Command, args []string) error {

	filename := args[0]
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Error reading resource file: %w", err)
	}

	var content []byte
	if withContent {
		content = fileBytes
	}

	sha256Digest := hex.EncodeToString(policy.GenSHA256(fileBytes))

	rdName := filename
	if len(name) > 0 {
		rdName = name
	}

	annotations, err := readAnnotations(annotationsFile)
	if err != nil {
		return fmt.Errorf("Error reading annotations file: %w", err)
	}
	
	rd := &ita.ResourceDescriptor{
		Name: rdName,
		Uri: uri,
		Digest: map[string]string{"sha256": strings.ToLower(sha256Digest)},
		Content: content,
		DownloadLocation: downloadLocation,
		MediaType: mediaType,
		Annotations: annotations,
	}

	err = rd.Validate()
	if err != nil {
		return fmt.Errorf("Invalid resource descriptor: %w", err)
	}
	
	return fileio.WritePbToFile(rd, outFile)
}

func genRdForRemote(cmd *cobra.Command, args []string) error {

	remoteUri := args[0]

	var digestSet map[string]string
	if len(digest) > 0 {
		// the in-toto spec expects a hex-encoded string in DigestSets
		// https://github.com/in-toto/attestation/blob/main/spec/v1/digest_set.md
		_, err := hex.DecodeString(digest)
		if err != nil {
			return fmt.Errorf("Digest is not valid hex-encoded string: %w", err)
		}
		
		// we can assume that we have both variables set at this point
		digestSet = map[string]string{hashAlg: strings.ToLower(digest)}
	}

	annotations, err := readAnnotations(annotationsFile)
	if err != nil {
		return fmt.Errorf("Error reading annotations file: %w", err)
	}
	
	rd := &ita.ResourceDescriptor{
		Name: name,
		Uri: remoteUri,
		Digest: digestSet,
		DownloadLocation: downloadLocation,
		Annotations: annotations,
	}

	err = rd.Validate()
	if err != nil {
		return fmt.Errorf("Invalid resource descriptor: %w", err)
	}
	
	return fileio.WritePbToFile(rd, outFile)
}
