# SCAI Examples

This directory contains examples for SCAI use cases:

* [Binary attributes](./gcc-helloworld)
* [Build process attributes](./sbom+slsa)
* [Build environment attributes](./hermetic-evidence)
* [Build platform attributes](./secure-boot)

## Usage

Before running any example, make sure to follow the [setup instructions].

Each directory contains a script to run the example:
```bash
./run-example.sh
```

The resulting metadata will be stored in the respective `metadata/` directory.

:: warn :: The in-toto attestations generated for these examples are not
digitally signed, and are only provided for demo purposes.

[setup instructions]: ../docs/usage.md
