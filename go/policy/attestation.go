package policy

import(
	"github.com/adityasaky/in-toto-attestation-verifier/verifier"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/interpreter"
	ita "github.com/in-toto/attestation/go/v1"
)

func getAttestationCELEnv() (*cel.Env, error) {
	return cel.NewEnv(
		cel.Types(&ita.Statement{}),
		cel.Variable("subject", cel.ListType(cel.ObjectType("in_toto_attestation.v1.ResourceDescriptor"))),
		cel.Variable("predicateType", cel.StringType),
		cel.Variable("predicate", cel.ObjectType("google.protobuf.Struct")),
	)
}

func getAttestationActivation(statement *ita.Statement) (interpreter.Activation, error) {
	return interpreter.NewActivation(map[string]any{
		"type":          statement.Type,
		"subject":       statement.Subject,
		"predicateType": statement.PredicateType,
		"predicate":     statement.Predicate,
	})
}

func ApplyAttestationRules(statement *ita.Statement, rules []verifier.Constraint) error {

	env, err := getAttestationCELEnv()
	if err != nil {
		return err
	}

	input, err := getAttestationActivation(statement)
	if err != nil {
		return err
	}

	return applyRules(env, input, rules)
}