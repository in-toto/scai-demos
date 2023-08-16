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
	Short: "Generate a JSON-encoded in-toto Resource Descriptor from local file(s)",
	RunE:  genRdFromFile,
}

var (
	withContent          bool
	downloadLocation string
	mediaType        string
	outFile          string
)

func init() {
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

	rdCmd.PersistentFlags().StringVarP(
		&outFile,
		"out-file",
		"o",
		"",
		"Filename to write out the RD object",
	)
	rdCmd.MarkPersistentFlagRequired("out-file")

	rdCmd.AddCommand(rdFileCmd)
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
		return fmt.Errorf("Error reading resource file", err)
	}

	var content []byte
	if withContent {
		content = fileBytes
	}

	sha256Digest := hex.EncodeToString(genSHA256(fileBytes))
	
	rd := &ita.ResourceDescriptor{
		Name: filename,
		Digest: map[string]string{"sha256": strings.ToLower(sha256Digest)},
		Content: content,
		DownloadLocation: downloadLocation,
		MediaType: mediaType,
	}

	err = rd.Validate()
	if err != nil {
		return fmt.Errorf("Invalid resource descriptor", err)
	}
	
	return util.WritePbToFile(rd, outFile)
}
