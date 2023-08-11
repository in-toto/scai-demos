# Build environment attributes

This example shows how SCAI can be used to capture information as evidence for
other supply chain metadata. Specifically, this example captures a run-time
trace of a SLSA builder, and uses this log as evidence for the [hermetic]
requirement of the SLSA v0.1 spec.

The SLSA Provenance and strace log files used in this example were generated
using [this workflow].

## A negative test

The SLSA build environment used to generate the metadata in this
example is not configured to support hermetic builds, so the included log
should not be considered valid evidence for the claimed `IsHermeticBuild`
attribute in the SCAI attribute assertion.

## Generated SCAI attestation

```jsonc
{
    "_type": "https://in-toto.io/Statement/v1",
    "subject": [
        {
            "name": "pdo_client_wawaka",
            "digest": {
                "sha256": "9b151e8b47a372bb686a441349d981ebf38951d70c4e7bf4669672651da7d33e"
            }
        }
    ],
    "predicateType": "https://in-toto.io/attestation/scai/attribute-report/v0.2",
    "predicate": {
        "attributes": [
            {
                "attribute": "IsHermeticBuild",
                "target": {
                    "name": "pdo_client_wawaka.provenance.json",
                    "digest": {
                        "sha256": "d094023e47dadedb5591d70bf203b967daa1e0784b29135053e410e6627ab261"
                    },
                    "downloadLocation": "https://github.com/marcelamelara/private-data-objects/suites/15001861846/artifacts/856214822",
                    "mediaType": "application/vnd.in-toto.provenance+json"
                },
                "evidence": {
                    "name": "strace.log",
                    "digest": {
                        "sha256": "1fbd9b1a3b867657666edf7ac1f17a3cad6ff691364ab46848ca5f49720688a8"
                    },
                    "downloadLocation": "https://github.com/marcelamelara/private-data-objects/suites/15001861846/artifacts/856214824",
                    "mediaType": "text/plain",
                    "annotations": {
                        "expectedResult": "fail"
                    }
                }
            }
        ]
    }
}
```

[hermetic]: https://slsa.dev/spec/v0.1/requirements#hermetic
[this workflow]: https://github.com/marcelamelara/private-data-objects/blob/generate-swsc-build-metadata/.github/workflows/ci-slsa3-tracing.yaml
