# scai-gen Go CLI

## Setup

Assuming you have Go version 1.20 or higher installed, run:

```bash
go build
go install
```

## Usage

To generate an in-toto [Resource Descriptor] for a local file:

```bash
scai-gen rd file -o <out-file> <filename> 
```

Run `scai-gen help` for a full list of command-line options.

[Resource Descriptor]: https://github.com/in-toto/attestation/blob/main/spec/v1/resource_descriptor.md
