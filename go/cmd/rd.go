package cmd

import (
	"fmt"
	"os"

	itpb "github.com/in-toto/attestation/go/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"github.com/spf13/cobra"
)

var rdCmd = &cobra.Command{
	Use:   "rd",
	Short: "Generates an in-toto Resource Descriptor file",
	RunE: genResourceDesc,
}

var (
	name             string
	uri              string
	digest           bool
	content          bool
	downloadLocation string
	mediaType        string
	resourceFile     string
	outFile          string
)

func init() {
	rdCmd.Flags().StringVarP(
		&name,
		"name",
		"n",
		"",
		"The name of the resource in the RD",
	)

	rdCmd.Flags().StringVarP(
		&uri,
		"uri",
		"u",
		"",
		"The URI of the resource in the RD",
	)

	rdCmd.Flags().BoolVarP(
		&digest,
		"digest",
		"d",
		false,
		"Flag to generate the SHA256 hash of the resource",
	)

	rdCmd.Flags().BoolVarP(
		&content,
		"content",
		"c",
		false,
		"Flag to include the content of the resource",
	)

	rdCmd.Flags().StringVarP(
		&downloadLocation,
		"download-location",
		"l",
		"",
		"The download location of the resource (if different from the URI)",
	)

	rdCmd.Flags().StringVarP(
		&mediaType,
		"media-type",
		"t",
		"",
		"The media type of the resource",
	)

	rdCmd.Flags().StringVarP(
		&resourceFile,
		"resource-file",
		"f",
		"",
		"Filename of the resource to be described",
	)

	rdCmd.Flags().StringVarP(
		&outFile,
		"out-file",
		"o",
		"",
		"Filename to write out the RD object",
	)
	rdCmd.MarkFlagRequired("out-file")
}

func genResourceDesc(cmd *cobra.Command, args []string) error {

	if len(name) == 0 && len(uri) == 0 && !digest {
		return fmt.Errorf("Need at least one of name, URI or digest for a valid in-toto resource descriptor")
	}

	/*
	contents := nil
	if len(resourceFile) > 0 {
		contents, err := os.ReadFile(resourceFile)
	}

	if err != nil {
		fmt.Println("Error reading file", err)
		return err
	}
	*/

	rd := &itpb.ResourceDescriptor{
		Name: name,
		Uri: uri,
		DownloadLocation: downloadLocation,
		MediaType: mediaType,
	}

	err := rd.Validate()
	if err != nil {
		return fmt.Errorf("Invalid resource descriptor", err)
	}

	rdBytes, err := protojson.Marshal(rd)
	if err != nil {
		return err
	}

	return os.WriteFile(outFile, rdBytes, 0644)
}
