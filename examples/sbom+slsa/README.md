# SBOM + SLSA example

This example shows how SCAI can be used to bind multiple pieces of metadata
(in this case an SPDX SBOM and a SLSA Provenance attestation) to capture
multiple attributes about an artifact's build process or supply chain.

To showcase further integration with the SW supply chain ecosystem, the
generated attestation in this example matches what an attestation to a [GUAC]
[evidence trees] for an SBOM and SLSA Provenance information might be.

The SPDX and SLSA Provenance files used in this example were generated using
[this workflow].

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
                "attribute": "HasSBOM",
                "evidence": {
                    "mediaType": "application/spdx+json",
                    "digest": {
                        "sha256": "911d4365b61ba7ace55f7333b2c638caca4b811ee73da5beb28b9ecbbd22ca78"
                    },
                    "downloadLocation": "https://github.com/marcelamelara/private-data-objects/suites/14359811861/artifacts/808758122"
                }
            },
            {
                "attribute": "HasSLSA",
                "evidence": {
                    "mediaType": "application/x.dsse+jsonl",
                    "digest": {
                        "sha256": "ea4d1e56e739f26a451c095b9fb40a353b3e73ea1778fdddafe13562e81bd745"
                    },
                    "downloadLocation": "https://github.com/marcelamelara/private-data-objects/suites/14359811861/artifacts/808758121"
                }
            }
        ]
    }
}
```

[GUAC]: https://github.com/guacsec/guac
[evidence trees]: https://docs.guac.sh/graphql/#the-guac-evidence-trees
[this workflow]: https://github.com/marcelamelara/private-data-objects/blob/generate-swsc-build-metadata/.github/workflows/ci-swsc.yaml
