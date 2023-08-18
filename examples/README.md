# SCAI Examples

This directory contains examples for SCAI use cases:

* [Binary attributes](./gcc-helloworld)
* [Build process attributes](./sbom+slsa)
* [Build platform attributes](./secure-boot)

## Usage

Before running any example, make sure to follow the [setup instructions].

Each directory contains a script to run the example:
```bash
./run-example.sh
```

The resulting metadata will be stored in the respective `metadata/` directory.

[setup instructions]: ../docs/usage.md
