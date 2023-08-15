# Dependency vulnerability attributes

This example shows how SCAI can be used to make evidence-based assertions
about third-party dependencies. Specifically, this example captures the
results of a vulnerability scan (Snyk) as evidence for the claim that no
known vulnerabilities were found in the dependencies of source code stored
in a particular git repository.

The Snyk results file used in this example was generated using a private
workflow.

## Generated SCAI attestation

```jsonc
{
    "_type":"https://in-toto.io/Statement/v1",
    "subject":[
        {
            "uri":"https://github.com/IntelLabs/supply-chain-attribute-integrity",
            "digest": {
                "sha1":"caf62c3caf2ab1228fb546f149f26df23c308802"
            }
        }
    ],
    "predicateType":"https://in-toto.io/attestation/scai/attribute-report/v0.2",
    "predicate":{
        "attributes":[
            {
                "attribute":"NO KNOWN VULNERABILITIES",
                "evidence":{
                    "name":"snyk-results.txt",
                    "digest":{
                        "sha256":"6db3ee1d8bdf7ea064281ecfaf999f9eb5749e4647d2a59e62eaa1ea8fcf6837"
                    },
                    "downloadLocation":"https://scans.example.com/results/snyk",
                    "mediaType":"text/plain"
                }
            }
        ]
    }
}
```
