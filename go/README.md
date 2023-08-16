# scai-gen Go CLI

## Setup

Assuming you have Go version 1.20 or higher installed, run:

```bash
go build
go install
```

## Usage

scai-gen can be used to generate JSON encoded in-toto [Resource Descriptor]s,
SCAI [Attribute Assertions], and SCAI [Attribute Reports].


### Generate an in-toto Resource Descriptor

Local file:

```bash
scai-gen rd file -o <out-file> [-n <name>] [-u <URI>] [-l <download location>] [-t <media type>] <filename> 
```

Remote resource or service:

```bash
scai-gen rd remote -o <out-file> [-a <hash algorithm> -d <digest>] [-n <name>] <resource URI> 
```

Run `scai-gen help` for a full list of command-line options.

[Resource Descriptor]: https://github.com/in-toto/attestation/blob/main/spec/v1/resource_descriptor.md
