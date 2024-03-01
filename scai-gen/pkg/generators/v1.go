package generators

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/in-toto/scai-demos/scai-gen/pkg/policy"

	ita "github.com/in-toto/attestation/go/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

// Generates an in-toto Attestation Framework v1 ResourceDescriptor for a local file, including its digest (default sha256).
// Throws an error if the resulting ResourceDescriptor does not meet the spec.
func NewRdForFile(filename, name, uri, hashAlg string, withContent bool, mediaType, downloadLocation string, annotations *structpb.Struct) (*ita.ResourceDescriptor, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading resource file: %w", err)
	}

	var content []byte
	if withContent {
		content = fileBytes
	}

	var digest string
	var alg string
	if hashAlg == "sha256" || hashAlg == "" {
		digest = hex.EncodeToString(policy.GenSHA256(fileBytes))
		alg = "sha256"
	} else {
		return nil, fmt.Errorf("hash algorithm %s not supported", hashAlg)
	}

	rdName := filename
	if len(name) > 0 {
		rdName = name
	}

	rd := &ita.ResourceDescriptor{
		Name:             rdName,
		Uri:              uri,
		Digest:           map[string]string{alg: strings.ToLower(digest)},
		Content:          content,
		DownloadLocation: downloadLocation,
		MediaType:        mediaType,
		Annotations:      annotations,
	}

	err = rd.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid resource descriptor: %w", err)
	}

	return rd, nil
}

// Generates an in-toto Attestation Framework v1 ResourceDescriptor for a remote resource identified by a name or URI).
// Does not check if the URI resolves to a valid remote location.
// Throws an error if the resulting ResourceDescriptor does not meet the spec.
func NewRdForRemote(name, uri, hashAlg, digest, downloadLocation string, annotations *structpb.Struct) (*ita.ResourceDescriptor, error) {
	digestSet := make(map[string]string)
	if len(hashAlg) > 0 && len(digest) > 0 {
		digestSet = map[string]string{hashAlg: strings.ToLower(digest)}
	}

	rd := &ita.ResourceDescriptor{
		Name:             name,
		Uri:              uri,
		Digest:           digestSet,
		DownloadLocation: downloadLocation,
		Annotations:      annotations,
	}

	err := rd.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid resource descriptor: %w", err)
	}

	return rd, nil
}

// Generates an in-toto Attestation Framework v1 Statement including a given predicate.
// Throws an error if the resulting Statement does not meet the spec.
func NewStatement(subjects []*ita.ResourceDescriptor, predicateType string, predicate *structpb.Struct) (*ita.Statement, error) {
	statement := &ita.Statement{
		Type:          ita.StatementTypeUri,
		Subject:       subjects,
		PredicateType: predicateType,
		Predicate:     predicate,
	}

	err := statement.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid in-toto Statement: %w", err)
	}

	return statement, nil
}
