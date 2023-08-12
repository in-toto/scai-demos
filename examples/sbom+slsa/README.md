# SBOM + SLSA example

This example shows how SCAI can be used to bind multiple pieces of metadata
(in this case an SPDX SBOM and a SLSA Provenance attestation) to capture
multiple attributes about an artifact's build process or supply chain.

The SPDX and SLSA Provenance files used in this example were generated using
[this workflow].

[this workflow]: https://github.com/marcelamelara/private-data-objects/blob/generate-swsc-build-metadata/.github/workflows/ci-swsc.yaml
