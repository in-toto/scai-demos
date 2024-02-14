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

func NewRdForFile(filename string, name string, uri string, hashAlg string, withContent bool, mediaType string, downloadLocation string, annotations *structpb.Struct) (*ita.ResourceDescriptor, error) {
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

func NewRdForRemote(name string, uri string, hashAlg string, digest string, downloadLocation string, annotations *structpb.Struct) (*ita.ResourceDescriptor, error) {
	digestSet := make(map[string]string)
	if len(digest) > 0 {
		// the in-toto spec expects a hex-encoded string in DigestSets
		// https://github.com/in-toto/attestation/blob/main/spec/v1/digest_set.md
		_, err := hex.DecodeString(digest)
		if err != nil {
			return nil, fmt.Errorf("digest is not valid hex-encoded string: %w", err)
		}

		// we can assume that we have both variables set at this point
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
