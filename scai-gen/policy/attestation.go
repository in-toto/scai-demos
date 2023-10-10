package policy

import(
	"github.com/adityasaky/in-toto-attestation-verifier/verifier"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/interpreter"
	ita "github.com/in-toto/attestation/go/v1"
	scai "github.com/in-toto/attestation/go/predicates/scai/v0"
)

func getAttestationCELEnv() (*cel.Env, error) {
	return cel.NewEnv(
		cel.Types(&ita.Statement{}),
		cel.Variable("subject", cel.ListType(cel.ObjectType("in_toto_attestation.v1.ResourceDescriptor"))),
		cel.Variable("predicateType", cel.StringType),
		cel.Variable("predicate", cel.ObjectType("google.protobuf.Struct")),
		cel.Types(&scai.AttributeAssertion{}),
		cel.Variable("assertion", cel.ObjectType("in_toto_attestation.predicates.scai.v0.AttributeAssertion")),
	)
}

func getAttestationActivation(statement *ita.Statement, attrAssertion *scai.AttributeAssertion) (interpreter.Activation, error) {
	return interpreter.NewActivation(map[string]any{
		"type":          statement.Type,
		"subject":       statement.Subject,
		"predicateType": statement.PredicateType,
		"predicate":     statement.Predicate,
		"assertion":     attrAssertion,
	})
}

func ApplyAttestationRules(statement *ita.Statement, attrAssertion *scai.AttributeAssertion, rules []verifier.Constraint) error {

	env, err := getAttestationCELEnv()
	if err != nil {
		return err
	}

	input, err := getAttestationActivation(statement, attrAssertion)
	if err != nil {
		return err
	}

	return applyRules(env, input, rules)
}