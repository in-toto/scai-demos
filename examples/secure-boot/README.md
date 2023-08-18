# Secure boot-attested builder example

This example shows how to expose a hardware platform attestation for a SW
supply chain component, like a build system. Using TPM 2.0-based secure boot
as an example, this proof of concept demonstrates a way for both build service
providers and their users to check the integrity of the build system itself.

For consistency in language and standards conformance, this example uses
terminology and concepts from the IETF [RATS] (RFC 9334) rfc standard.

## Attestation flow

To use secure boot attestations to check build system integrity, a two-stage
attetation process is required. First, a good-known reference value for the
state of the build platform is needed, against which the secure boot
attestation (i.e., the evidence) can be later compared. In this example, we
first generate such a TPM reference value attestation.

### TPM reference value SCAI attestation

```jsonc
{
    "_type": "https://in-toto.io/Statement/v1",
    "subject": [
        {
            "digest": {
                "sha256": "7f13e3ce0c34086f3a4e4aaccab8b42eddd0a8e8aa5efe817905f93522e17739"
            },
            "content": "eyJoYXNfc2VjdXJlYm9vdCI6IHRydWUsICJzY3J0b...",
            "mediaType": "application/json"
        }
    ],
    "predicateType": "https://in-toto.io/attestation/scai/attribute-report/v0.2",
    "predicate": {
        "attributes": [
            {
                "attribute": "TPM2.0_REFERENCE_STATE",
                "target": {
                    "name": "localhost",
                    "uri": "http://127.0.0.1"
                }
            }
        ]
    }
}
```

In this case, this TPM reference value (obtained using the [Keylime]
framework for TPM 2.0 attestation) is itself the subject artifact of the
attestation. The second stage of build system integrity checking then
requires the secure boot quote, which together with the reference value can
be used to check that the build system used to build binary (such as
hello-world) has tamper-evident properties.

### Build platform attestation

```jsonc
{
    "_type": "https://in-toto.io/Statement/v1",
    "subject": [
        {
            "name": "hello-world",
            "digest": {
                "sha256": "4ad979bf2f60aa07011b974094a415e05d238d7507f8c65568a76c6c291a6b62"
            }
        }
    ],
    "predicateType": "https://in-toto.io/attestation/scai/attribute-report/v0.2",
    "predicate": {
        "attributes": [
            {
                "attribute": "TAMPER-EVIDENT BUILD PLATFORM",
                "target": {
                    "uri": "http://127.0.0.1",
                    "name": "localhost"
                },
                "evidence": {
                    "name": "secure boot quote",
                    "digest": {
                        "sha256": "0dc21d96a4f6f86f015b498f979bb1ce0daa1cc19903d9d1109a0ef4819f6acd"
                    },
                    "mediaType": "application/json"
                }
            }
        ]
    }
}
```

For this second stage of build system attestation, the secure boot quote
provided by the TPM is the evidence for the tamper-evidence claim.
A relying party seeking to validate the build system for the hello-world
binary would then require both the TPM reference value and secure boot quote
attestations.

[Keylime]: https://keylime.dev/
[RATS]: https://datatracker.ietf.org/doc/rfc9334/
