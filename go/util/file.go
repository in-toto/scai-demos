package util

import(
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
