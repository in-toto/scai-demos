package generators

import (
	"fmt"

	scai "github.com/in-toto/attestation/go/predicates/scai/v0"
	ita "github.com/in-toto/attestation/go/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

func NewSCAIAssertion(attribute string, target *ita.ResourceDescriptor, conditions *structpb.Struct, evidence *ita.ResourceDescriptor) (*scai.AttributeAssertion, error) {
	aa := &scai.AttributeAssertion{
		Attribute:  attribute,
		Target:     target,
		Conditions: conditions,
		Evidence:   evidence,
	}

	err := aa.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid SCAI attribute assertion: %w", err)
	}

	return aa, nil
}

func NewSCAIReport(attrAssertions []*scai.AttributeAssertion, producer *ita.ResourceDescriptor) (*scai.AttributeReport, error) {
	ar := &scai.AttributeReport{
		Attributes: attrAssertions,
		Producer:   producer,
	}

	err := ar.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid SCAI attribute report: %w", err)
	}

	return ar, nil
}
