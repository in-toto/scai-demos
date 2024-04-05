// adapted from https://github.com/slsa-framework/slsa-github-generator/blob/main/signing/sigstore/fulcio.go
// and https://github.com/slsa-framework/slsa-github-generator/blob/main/internal/builders/generic/attest.go
package cmd

import (
	"bytes"
	"context"
	"fmt"

	"github.com/in-toto/scai-demos/scai-gen/pkg/fileio"

	ita "github.com/in-toto/attestation/go/v1"
	"github.com/sigstore/cosign/v2/cmd/cosign/cli/fulcio"
	"github.com/sigstore/cosign/v2/cmd/cosign/cli/options"
	"github.com/sigstore/cosign/v2/cmd/cosign/cli/sign"
	"github.com/sigstore/cosign/v2/pkg/providers"
	"github.com/sigstore/sigstore/pkg/signature/dsse"
	"github.com/slsa-framework/slsa-github-generator/signing/envelope"
	"github.com/slsa-framework/slsa-github-generator/signing/sigstore"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

var sigstoreCmd = &cobra.Command{
	Use:   "sigstore",
	Args:  cobra.ExactArgs(1),
	Short: "Use Sigstore signing to create a signed DSSE for an in-toto Statements",
	RunE:  signWithSigstore,
}

// attestation is a signed attestation.
type attestation struct {
	cert []byte
	att  []byte
}

// Bytes returns the signed attestation as an encoded DSSE JSON envelope.
func (a *attestation) Bytes() []byte {
	return a.att
}

// Cert returns the certificate used to sign the attestation.
func (a *attestation) Cert() []byte {
	return a.cert
}

func init() {
	sigstoreCmd.Flags().StringVarP(
		&outFile,
		"out-file",
		"o",
		"",
		"Filename to write out the signed DSSE object",
	)
	sigstoreCmd.MarkFlagRequired("out-file") //nolint:errcheck
}

func getNewFulcioSigner(ctx context.Context) (*fulcio.Signer, error) {
	ko := options.KeyOpts{
		OIDCIssuer:   options.DefaultOIDCIssuerURL,
		OIDCClientID: "sigstore",
		FulcioURL:    options.DefaultFulcioURL,
	}

	sv, err := sign.SignerFromKeyOpts(ctx, "", "", ko)
	if err != nil {
		return nil, fmt.Errorf("getting Fulcio signer: %w", err)
	}

	return fulcio.NewSigner(ctx, ko, sv)
}

func signWithSigstore(cmd *cobra.Command, args []string) error {
	fmt.Println("EXPERIMENTAL FEATURE. DO NOT USE IN PRODUCTION.")

	// want to make sure the DSSE is a JSON file
	if !fileio.HasJSONExt(outFile) {
		return fmt.Errorf("expected a .json extension for the generated DSSE file %s", outFile)
	}

	statementFile := args[0]
	statement := &ita.Statement{}
	err := fileio.ReadPbFromFile(statementFile, statement)
	if err != nil {
		return err
	}

	// will only sign valid Statements
	err = statement.Validate()
	if err != nil {
		return fmt.Errorf("invalid in-toto Statement: %w", err)
	}

	ctx := cmd.Context()

	// Get Fulcio signer
	if !providers.Enabled(ctx) {
		return fmt.Errorf("no auth provider is enabled. Are you running outside of Github Actions?")
	}

	attBytes, err := protojson.Marshal(statement)
	if err != nil {
		return fmt.Errorf("error marshalling Statement: %w", err)
	}

	k, err := getNewFulcioSigner(ctx)
	if err != nil {
		return fmt.Errorf("error creating Fulcio signer: %w", err)
	}

	dsseSigner := dsse.WrapSigner(k, "application/vnd.in-toto")
	signedAtt, err := dsseSigner.SignMessage(bytes.NewReader(attBytes))
	if err != nil {
		return fmt.Errorf("error signing DSSE: %w", err)
	}

	// Add certificate to envelope. This is needed for
	// Rekor compatibility.
	signedAttWithCert, err := envelope.AddCertToEnvelope(signedAtt, k.Cert)
	if err != nil {
		return fmt.Errorf("error adding Fulcio certificate to DSSE: %w", err)
	}

	tlog := sigstore.NewDefaultRekor()
	_, err = tlog.Upload(ctx, &attestation{
		att:  signedAttWithCert,
		cert: k.Cert,
	})
	if err != nil {
		return fmt.Errorf("error uploading signed DSSE to public Rekor log: %w", err)
	}

	return fileio.WriteDSSEToFile(signedAtt, outFile)
}
