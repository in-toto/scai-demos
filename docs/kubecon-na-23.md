# KubeCon NA '23 Demo

As part of the [in-toto Maintainer Track talk] at KubeCon NA '23, we present
a demo of the in-toto Attestation Framework, SCAI, and the in-toto Attestation
Verifier.

## Demo Setup

The overall flow implemented in the demo is as follows:

<img src="./images/scai-kubeconNA23-demo.png" alt="in-toto demo flow" width="600">

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

### Additional Tools

This demo makes use of the following additional tools:

* in-toto [attestation-verifier]
* [Anchore SBOM generator] GitHub Action
* [SLSA generic Provenance generator] GitHub Action
* [strace] Linux syscall tracer

[Anchore SBOM generator]: https://github.com/anchore/sbom-action
[attestation-verifier]: https://github.com/in-toto/attestation-verifier
[demo workflow]: https://github.com/marcelamelara/private-data-objects/blob/kubeconNA23-intoto-demo/.github/workflows/kubeconNA23-intoto-demo.yml
[in-toto Maintainer Track talk]: https://kccncna2023.sched.com/event/1R2mx
[SLSA generic Provenance generator]: https://github.com/slsa-framework/slsa-github-generator
[SLSA Provenance]: https://github.com/in-toto/attestation/blob/main/spec/predicates/provenance.md
[SCAI Attribute Report]: https://github.com/in-toto/attestation/blob/main/spec/predicates/scai.md
[scai-gen GitHub Actions]: https://github.com/in-toto/scai-demos/tree/main/.github/actions
[strace]: https://strace.io/