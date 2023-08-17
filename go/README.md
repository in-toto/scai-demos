# scai-gen Go CLI

This package provides a Go CLI for generating in-toto compatible [SCAI] metadata.

## Setup

Assuming you have Go version 1.20 or higher installed, run:

```bash
go build
go install
```

## Usage

scai-gen can be used to generate JSON encoded in-toto [Resource Descriptors],
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

### Generate a SCAI Attribute Assertion

```bash
scai-gen assert -o <out-file> [-e <evidence RD filename>] <attribute> 
```

Run `scai-gen assert help` for a full list of command-line options.

### Generate a SCAI Attribute Report

```bash
scai-gen report -o <out-file> [-e <evidence RD filename>] <attribute assertion file1> [<attribute assertion file2> ...]
```

Run `scai-gen report help` for a full list of command-line options.

[Attribute Assertions]: https://github.com/in-toto/attestation/blob/main/protos/in_toto_attestation/predicates/scai/v0/scai.proto#L16
[Attribute Reports]: https://github.com/in-toto/attestation/blob/main/protos/in_toto_attestation/predicates/scai/v0/scai.proto#L28
[Resource Descriptors]: https://github.com/in-toto/attestation/blob/main/spec/v1/resource_descriptor.md
[SCAI]: https://github.com/in-toto/attestation/blob/main/spec/predicates/scai.md
