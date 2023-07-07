# Introduction to SCAI

The Software Supply Chain Attribute Integrity, or SCAI (pronounced "sky"),
specification is a data format for claims about functional attributes and
integrity about a software artifact and its supply chain, along with any
evidence for the claim.

That is, in addition to code attributes, SCAI can capture metadata about any
layer of a software stack, from the tool that built an artifact, to
information about the software system or the hardware platform that ran a
given supply chain step.

SCAI data can be associated with executable binaries, statically- or
dynamically-linked libraries, software packages, container images,
software toolchains, and compute environments.

As such, SCAI is intended to be implemented as part of an existing software
supply chain attestation framework by software development tools or services
(e.g., builders, CI/CD pipelines, software analysis tools) seeking to
capture more granular information about the attributes and behavior of the
software artifacts they produce. That is, SCAI assumes that implementers will
have appropriate processes and tooling in place for capturing other types of
software supply chain metadata, which can be extended to add support for SCAI.

## Integration with in-toto

SCAI integrates with the [in-toto Attestation Framework] as an attestation
predicate that can be generated alongside any other supply chain metadata,
such as SBOM or SLSA Provenance.

[in-toto Attestation Framework]: https://github.com/in-toto/attestation