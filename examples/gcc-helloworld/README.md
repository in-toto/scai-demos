# Binary attributes example

This example shows how to capture fine-grained binary properties during
the compilation of a hello-world program based on the gcc flags used for a
given build.

## Generated SCAI attestation

```jsonc
{
    "_type": "https://in-toto.io/Statement/v1",
    "subject": [
        {
            "name": "hello-world",
            "digest": {
                "sha256": "0b222de8bcb1ea30807f1a4d733108e96de4512b689da8c5f371ac8a572e9271"
            }
        }
    ],
    "predicateType": "https://in-toto.io/attestation/scai/attribute-report/v0.2",
    "predicate": {
        "attributes": [
            {
                "attribute": "HAS_STACK_PROTECTION",
                "conditions": {
                    "flags": "-fstack-protector*"
                }
            }
        ],
        "producer": {
            "name": "gcc (Ubuntu 11.4.0-1ubuntu1~22.04) 11.4.0",
            "digest": {
                "sha256": "fb60c5945ce785d0cb2ef0303dac5249f25bd0d0324317a0734aae6aa24f609b"
            },
            "uri": "file://usr/bin/gcc",
            "annotations": {
                "command": "gcc -fstack-protector -o hello-world hello-world.c"
            }
        }
    }
}
```
