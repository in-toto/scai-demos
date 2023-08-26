package policy

import(
	"fmt"
	
	"github.com/adityasaky/in-toto-attestation-verifier/verifier"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/interpreter"
	scai "github.com/in-toto/attestation/go/predicates/scai/v0"
	//"google.golang.org/protobuf/types/known/structpb"
)

func getPlaintextCELEnv() (*cel.Env, error) {
	return cel.NewEnv(
		cel.Types(&scai.AttributeAssertion{}),
		cel.Variable("text", cel.StringType),
		cel.Variable("assertion", cel.ObjectType("in_toto_attestation.predicates.scai.v0.AttributeAssertion")),
	)
}

func getPlaintextActivation(text string, attrAssertion *scai.AttributeAssertion) (interpreter.Activation, error) {
	return interpreter.NewActivation(map[string]any{
		"text": text,
		"assertion": attrAssertion,
	})
}

func ApplyPlaintextRules(text string, attrAssertion *scai.AttributeAssertion, rules []verifier.Constraint) error {

	env, err := getPlaintextCELEnv()
	if err != nil {
		return fmt.Errorf("Failed to init CEL env: %w", err)
	}

	input, err := getPlaintextActivation(text, attrAssertion)
	if err != nil {
		return fmt.Errorf("Failed to get CEL activation: %w", err)
	}

	return applyRules(env, input, rules)
}