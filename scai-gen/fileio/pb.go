package fileio

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WritePbToFile(pb proto.Message, outFile string) error {
	pbBytes, err := protojson.Marshal(pb)
	if err != nil {
		return err
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
