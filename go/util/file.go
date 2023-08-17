package util

import(
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func WritePbToFile(pb proto.Message, outFile string) error {
	pbBytes, err := protojson.Marshal(pb)
	if err != nil {
		return err
	}

	return os.WriteFile(outFile, pbBytes, 0644)
}

func ReadPbFromFile(filename string, pb proto.Message) error {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Error reading file: %w", err)
	}

	err = protojson.Unmarshal(fileBytes, pb)
	if err != nil {
		return fmt.Errorf("Error unmarshalling protobuf: %w", err)
	}

	return nil
}