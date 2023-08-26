# scai-gen Go CLI

This package provides a Go CLI for generating in-toto compatible [SCAI]
metadata. We assume a minimal Ubuntu 20.04+ platform.

## Setup

First, install Go version 1.20 or higher following the
[Go installation instructions](https://go.dev/doc/install), as well as
additional dependencies:

```
sudo apt install build-essential
```

Then, install the scai-gen Go module from this repo's root directory:

```bash
make go-mod
```

## Usage

scai-gen can be used to generate JSON encoded in-toto [Resource Descriptors],
SCAI [Attribute Assertions], and SCAI [Attribute Reports].

scai-gen also provides a feature for checking [DSSE]-signed in-toto attestations
against an in-toto [Layout] or a SCAI [evidence policy]. Examples can be found
in the [layouts](../layouts) and [policies](../policies) directories.

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

### Check the in-toto Layout for SCAI attestations

```bash
scai-gen check layout -l <layout-file> <SCAI attestation file1> [<SCAI attestation file2> ...]
```

### Check SCAI attestation against an evidence policy

```bash
scai-gen check evidence -p <policy-file> -e <evidence-directory> <SCAI attestation file>
```

Run `scai-gen check help` for a full list of command-line options.

[Attribute Assertions]: https://github.com/in-toto/attestation/blob/main/protos/in_toto_attestation/predicates/scai/v0/scai.proto#L16
[Attribute Reports]: https://github.com/in-toto/attestation/blob/main/protos/in_toto_attestation/predicates/scai/v0/scai.proto#L28
[DSSE]: https://github.com/in-toto/attestation/blob/main/spec/v1/envelope.md
[Layout]: https://github.com/in-toto/ITE/tree/master/ITE/10
[Resource Descriptors]: https://github.com/in-toto/attestation/blob/main/spec/v1/resource_descriptor.md
[SCAI]: https://github.com/in-toto/attestation/blob/main/spec/predicates/scai.md
[evidence policy]: ./policy/checks.go#L15
