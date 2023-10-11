package policy

import(
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	
	"github.com/in-toto/attestation-verifier/verifier"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/interpreter"
)

type SCAIEvidencePolicy struct {
	AttestationID string                 `yaml:"attestationID"`
	Inspections   []*verifier.Inspection `yaml:"inspections"`
}

func GenSHA256(bytes []byte) []byte {
	h := sha256.New()
	h.Write(bytes)
	return h.Sum(nil)
}

func MatchDigest(hexDigest string, blob []byte) bool {

	digest := GenSHA256(blob)

	decoded, err := hex.DecodeString(hexDigest)
	if err != nil {
		fmt.Println("Problem decoding hex-encoded digest to match")
		return false
	}

	return bytes.Equal(decoded, digest)
}

func applyRules(env *cel.Env, input interpreter.Activation, rules []verifier.Constraint) error {

	for _, r := range rules {

		ast, issues := env.Compile(r.Rule)
		if issues != nil && issues.Err() != nil {
			return fmt.Errorf("CEL compilation issues: %w", issues.Err())
		}

		prog, err := env.Program(ast)
		if err != nil {
			return fmt.Errorf("CEL program error: %w", err)
		}

		out, _, err := prog.Eval(input)
		if err != nil {
			if strings.Contains(err.Error(), "no such attribute") || strings.Contains(err.Error(), "no such key") && r.AllowIfNoClaim {
				continue
			}
			return err
		}
		switch result := out.Value().(type) {
		case bool:
			if !result {
				if !r.Warn {
					return fmt.Errorf("policy check failed for rule '%s'", r.Rule)
				}
				fmt.Println("Rule", r.Rule, "failed.")
			}
		case error:
			return fmt.Errorf("CEL error: %w", result)
		}
	}

	return nil
}
