// adapted from https://github.com/slsa-framework/slsa-github-generator/blob/main/signing/sigstore/fulcio.go
// and https://github.com/slsa-framework/slsa-github-generator/blob/main/internal/builders/generic/attest.go
package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rekorCmd = &cobra.Command{
	Use:   "rekor",
	Args:  cobra.ExactArgs(1),
	Short: "Parses a Rekor log entry to extract info needed to verify signed in-toto Attestations",
	RunE:  parseRekorEntry,
}

func parseRekorEntry(cmd *cobra.Command, args []string) error {
	fmt.Println("EXPERIMENTAL FEATURE. DO NOT USE IN PRODUCTION.")

	entryFile := args[0]
	readFile, err := os.Open(entryFile)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines [][]byte

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Bytes())
	}

	readFile.Close()

	for _, line := range fileLines {
		lineStr := string(line)
		if strings.Contains(lineStr, "publicKey") {
			pkB64Raw := strings.TrimPrefix(lineStr, "    \"publicKey\": ")
			pkB64 := strings.Trim(pkB64Raw, "\"")

			pkPem, err := base64.StdEncoding.DecodeString(pkB64)
			if err != nil {
				return fmt.Errorf("error decoding base64-encoded public key: %w", err)
			}

			// lazy
			fmt.Println(string(pkPem))
			return nil
		}
	}

	return nil
}
