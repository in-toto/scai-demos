package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"scai-gen/util"

	ita "github.com/in-toto/attestation/go/v1"
	"github.com/spf13/cobra"
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
)

func init() {
	rdCmd.AddCommand(rdFileCmd)
	rdCmd.AddCommand(rdRemoteCmd)
}

func init() {
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
		"a",
		"",
		"The hash algorithm used to compute the digest associated with the remote resource",
	)
	rdRemoteCmd.MarkFlagsRequiredTogether("digest", "hash-alg")
}

func genSHA256(bytes []byte) []byte {
	h := sha256.New()
	h.Write(bytes)
	return h.Sum(nil)
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

	sha256Digest := hex.EncodeToString(genSHA256(fileBytes))
	
	rd := &ita.ResourceDescriptor{
		Name: filename,
		Uri: uri,
		Digest: map[string]string{"sha256": strings.ToLower(sha256Digest)},
		Content: content,
		DownloadLocation: downloadLocation,
		MediaType: mediaType,
	}

	err = rd.Validate()
	if err != nil {
		return fmt.Errorf("Invalid resource descriptor: %w", err)
	}
	
	return util.WritePbToFile(rd, outFile)
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
	
	rd := &ita.ResourceDescriptor{
		Name: name,
		Uri: remoteUri,
		Digest: digestSet,
		DownloadLocation: downloadLocation,
	}

	err := rd.Validate()
	if err != nil {
		return fmt.Errorf("Invalid resource descriptor: %w", err)
	}
	
	return util.WritePbToFile(rd, outFile)
}
