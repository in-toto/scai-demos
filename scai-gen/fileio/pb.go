package fileio

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WritePbToFile(pb proto.Message, outFile string, pretty bool) error {
	opt := &protojson.MarshalOptions{}
	if pretty {
		opt.Multiline = true
		opt.Indent = "    "
		opt.EmitUnpopulated = false
	}

	pbBytes, err := opt.Marshal(pb)
	if err != nil {
		return err
	}

	// ensure the out directory exists
	if err = CreateOutDir(outFile); err != nil {
		return fmt.Errorf("error creating output directory for file %s: %w", outFile, err)
	}

	return os.WriteFile(outFile, pbBytes, 0644) //nolint:gosec
}

func ReadPbFromFile(filename string, pb proto.Message) error {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	err = protojson.Unmarshal(fileBytes, pb)
	if err != nil {
		return fmt.Errorf("error unmarshalling protobuf: %w", err)
	}

	return nil
}
