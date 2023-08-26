# SCAI Examples

This directory contains examples for SCAI use cases:

* [Binary attributes](./gcc-helloworld)
* [Build process attributes](./sbom+slsa)
* [Build environment attributes](./hermetic-evidence)
* [Build platform attributes](./secure-boot)
* [Dependency vulnerability attributes](./vuln-scan)

## Basic Usage

Before running any example, make sure to follow the [setup instructions].

Each directory contains a script to run the Python example:

```bash
./run-example.sh
```

And to run the Go example:

```bash
./run-go-example.sh
```

The resulting metadata will be stored in the respective `metadata/` directory.

:: warn :: The scai-generator CLI tools do not generate digitally signed in-toto
attestations. Any digitally signed attestations included in this repo are only
provided for demo purposes.

## End-to-End Examples

We provide one end-to-end example for a container build that showcases both
SCAI attestation generation, as well as verification of an ITE-10 in-toto layout
and SCAI policy checking.

Use the `examples/run-container-examples-e2e.sh` bash script to run the full
example.

[setup instructions]: ../docs/usage.md
