package fileio

import (
	"encoding/json"
	"fmt"
	"os"

	ita "github.com/in-toto/attestation/go/v1"
	"github.com/secure-systems-lab/go-securesystemslib/dsse"
	"google.golang.org/protobuf/encoding/protojson"
)

func ReadDSSEFile(path string) (*dsse.Envelope, error) {
	envBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	envelope := &dsse.Envelope{}
	if err := json.Unmarshal(envBytes, envelope); err != nil {
		return nil, err
	}

	return envelope, nil
}

func ReadStatementFromDSSEFile(path string) (*ita.Statement, error) {
	envelope, err := ReadDSSEFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read DSSE file %w", err)
	}

	stBytes, err := envelope.DecodeB64Payload()
	if err != nil {
		return nil, err
	}

	statement := &ita.Statement{}
	if err = protojson.Unmarshal(stBytes, statement); err != nil {
		return nil, err
	}

	return statement, nil
}

func WriteDSSEToFile(envBytes []byte, outFile string) error {
	return os.WriteFile(outFile, envBytes, 0644) //nolint:gosec
}