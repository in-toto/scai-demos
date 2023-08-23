# scai-generator Usage

We encourage you to gain a basic understanding of the [SCAI specification]
before using the CLI tools in this repo.

The general flow is to first generate any [ResourceDescriptors],
one or more [AttributeAssertions] and then generate a SCAI [Report].
The generated SCAI report document is a valid [in-toto Statement].

Note, that the CLI tools do not generate **signed** SCAI Reports or
in-toto attestations.

## CLI Usage

The SCAI CLI tools and examples have been tested on Ubuntu 20.04 or higher.

For information on how to use our CLI tools in [Python] or [Go] environments,
please refer to their instructions.

[in-toto Statement]: https://github.com/in-toto/attestation/blob/main/spec/v1/statement.md
[Resource Descriptors]: https://github.com/in-toto/attestation/blob/main/spec/v1/resource_descriptor.md
[Attribute Assertions]: https://github.com/in-toto/attestation/blob/main/protos/in_toto_attestation/predicates/scai/v0/scai.proto#L16
[Report]: https://github.com/in-toto/attestation/blob/main/protos/in_toto_attestation/predicates/scai/v0/scai.proto#L28
[SCAI specification]: https://github.com/in-toto/attestation/blob/main/spec/predicates/scai.md
[Go]: ../go/README.md
[Python]: ../python/README.md
