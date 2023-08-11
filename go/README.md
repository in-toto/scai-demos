# scai-gen Go CLI

## Setup

Assuming you have Go version 1.20 or higher installed, run:

```bash
cd cmd && go build
go install
```

## Usage

To generate an in-toto [Resource Descriptor]:

```bash
scai-gen rd -n <name> -u <uri> -d -o <out-file>
```

[Resource Descriptor]: https://github.com/in-toto/attestation/blob/main/spec/v1/resource_descriptor.md
