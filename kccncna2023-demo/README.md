# KubeCon + CloudNativeCon NA '23 Demo

As part of the [in-toto Maintainer Track talk] at KubeCon + CloudNativeCon NA
'23, we present a demo of the in-toto Attestation Framework, SCAI, and the
in-toto Attestation Verifier.

## Demo Setup

The overall flow implemented in the demo is as follows:

<img src="./images/intoto-kccncna2023-demo.png" alt="in-toto demo flow" width="600">

This demo setup is implemented using the [scai-gen GitHub Actions] in a Docker
container build [demo workflow] for the Hyperledger Labs Private Data Objects
project.

### Generated Attestations

This demo generates the follow _authenticated_ in-toto attestations:

* [SLSA Provenance] attestation for the container build
* [SCAI Attribute Report] attestation for additional integrity metadata about
the build

These two attestations are signed using cosign OIDC-based keyless signing,
and uploaded to the public Rekor log.

### Verified Policies

This demo verifies the following policies using the generated attestations:

* [in-toto Layout] checks that the expected attestations were generated for each step
of the demo workflow.
* [SCAI policy] checks the attested attributes against the evidence indicated in the
SCAI Attribute Report.

This verification flow is implemented in the [verification-flow.sh] script.

### Additional Tools

This demo makes use of the following additional tools:

* in-toto [attestation-verifier]
* [Anchore SBOM generator] GitHub Action
* [SLSA generic Provenance generator] GitHub Action
* [strace] Linux syscall tracer

[Anchore SBOM generator]: https://github.com/anchore/sbom-action
[attestation-verifier]: https://github.com/in-toto/attestation-verifier
[demo workflow]: https://github.com/marcelamelara/private-data-objects/blob/intoto-kccncna2023-demo/.github/workflows/intoto-kccncna2023-demo.yml
[in-toto Layout]: ./policies/layout.yml
[in-toto Maintainer Track talk]: https://kccncna2023.sched.com/event/1R2mx
[SLSA generic Provenance generator]: https://github.com/slsa-framework/slsa-github-generator
[SLSA Provenance]: https://github.com/in-toto/attestation/blob/v1.0.1/spec/predicates/provenance.md
[SCAI Attribute Report]: https://github.com/in-toto/attestation/v1.0.1/main/spec/predicates/scai.md
[SCAI policy]: ./policies/has-slsa.yml
[scai-gen GitHub Actions]: https://github.com/in-toto/scai-demos/tree/main/.github/actions
[strace]: https://strace.io/
[verification-flow.sh]: ./verification-flow.sh
