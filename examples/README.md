# SCAI Examples

This directory contains examples for SCAI use cases:

* [Binary attributes](./gcc-helloworld)
* [Build process attributes](./sbom+slsa)
* [Build environment attributes](./hermetic-evidence)
* [Build platform attributes](./secure-boot)
* [Dependency vulnerability attributes](./vuln-scan)

## Usage

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

:: warn :: The in-toto attestations generated for these examples are not
digitally signed, and are only provided for demo purposes.

[setup instructions]: ../docs/usage.md
