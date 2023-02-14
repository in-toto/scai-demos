# Software Supply Chain Attribute Integrity

The Software Supply Chain Attribute Integrity, or SCAI (pronounced "sky"), specification proposes a data
format for capturing functional attribute and integrity information about software artifacts and their supply
chain. SCAI data can be associated with executable binaries, statically- or dynamically-linked libraries,
software packages, container images, software toolchains, and compute environments.

As such, SCAI is intended to be implemented as part of an existing software supply chain attestation
framework by software development tools or services (e.g., builders, CI/CD pipelines, software analysis tools)
seeking to capture more granular information about the attributes and behavior of the software artifacts they
produce. That is, SCAI assumes that implementers will have appropriate processes and tooling in place for
capturing other types of software supply chain metadata, which can be extended to add support for SCAI.

## Specification

The [SCAI specification] is hosted under the [in-toto Attestation Framework]
as an attestation predicate.

This repo also provides [JSON schema](schema/) that can be used in conjunction
with other software supply chain metadata.

## Documentation

All documentation can be found under [docs/](docs/).

## Usage

The general flow is to first generate one or more Attribute
Assertions and then generate a SCAI Report. The
[examples](examples/) show
how SCAI metadata is generated in a few different use cases.

Note, that the CLI tools do not current generate **signed**
SCAI Reports.

#### CLI Environment Setup

To run the SCAI CLI tools and examples, the following packages
are required on a minimal Ubuntu system. We assume Ubuntu 20.04.

```
sudo apt install git python3 python3-dev python3-venv virtualenv build-essential
```

Then, set up the Python virtualenv for the SCAI CLI tools.

```
make -C cli
```

#### Basic CLI Invocation

```
cd cli
source scai-venv/bin/activate
```

To generate a basic attribute assertion:
```
./scai-attr-assertion -a <attribute string> -o <output filename> [-e <evidence filename>] [-c <conditions string>]
```

To generate a basic SCAI Report with in-toto Link metadata:
```
./scai-report -i <input artifact filenames> -a <attribute assertion filenames> -c <command to execute as string>
```

For a full list of CLI tool options, invoke with the `-h` option.

[SCAI specification]: https://github.com/in-toto/attestation/blob/main/spec/predicates/scai.md
[in-toto Attestation Framework]: https://github.com/in-toto/attestation/tree/main/spec
