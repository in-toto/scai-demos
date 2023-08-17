# Software Supply Chain Attribute Integrity (SCAI) Generator

The Software Supply Chain Attribute Integrity, or SCAI (pronounced "sky"),
specification is a succint data format for claims and evidence about
attributes and integrity about a software artifact and its supply chain.

For more details read our [intro doc] or the full [SCAI spec doc].

## In this repo

This repo provides CLI tools for automatically generating valid,
in-toto compatible SCAI metadata, as well as example use cases for SCAI in
[examples](examples/).

The official [SCAI specification] is hosted under the
[in-toto Attestation Framework] as an attestation predicate.

All documentation can be found under [docs/](docs/).

## Usage

Read the [usage doc] for instructions on setup and tool invocation
for Python and Go environments.

We encourage you to gain a basic understanding of the [SCAI specification]
before using the scai-generator CLI tools in this repo.

## Disclaimer

While the tools in this repo are conformant to the
[in-toto Attestation Framework], they do not generate **authenticated** SCAI
attestations. The example use cases in this repo are only provided for
illustrative purposes.

[in-toto Attestation Framework]: https://github.com/in-toto/attestation/tree/main/spec
[intro doc]: docs/intro.md
[usage doc]: docs/usage.md
[SCAI specification]: https://github.com/in-toto/attestation/blob/main/spec/predicates/scai.md
[SCAI spec doc]: https://arxiv.org/pdf/2210.05813.pdf
