# scai-generator Python CLI

## CLI Environment Setup

To run the SCAI CLI tools and examples, the following packages
are required on a minimal Ubuntu system. We assume Ubuntu 20.04 or higher.

```
sudo apt install git python3 python3-dev python3-venv virtualenv build-essential
```

Then, set up the Python virtualenv for the SCAI CLI tools from this
repo's root directory.

```
make VENVDIR=<dest dir> py-venv
```

## Basic CLI Invocation

```
source $VENVDIR/bin/activate
```

### Generate a ResourceDescriptor

To generate ResourceDescriptors for SCAI AttributeAssertion or Report fields:
```
scai-gen-resource-desc -o <output filename> [-n <resource name> -d -u <resource URI>] [-r <resource filename>] [-l <resource download location>] [-c] [-t <resource media type>]
```

**:warn: Note:** Because at least one of `name`, `uri` or `digest` fields
are required in [ResourceDescriptors], the tool will throw an error if
none of these options are passed in.

### Generate an AttributeAssertion

To generate AttributeAssertions for a SCAI Report:
```
scai-attr-assertion -a <attribute string> -o <output filename> [-t <target filename>] [-e <evidence filename>] [-c <conditions filename>]
```

**:warn: Note:** Since the `conditions` field in [AttributeAssertions] can
be an arbitrary JSON object, the tool assumes that this object has been
written to a file beforehand.

### Generate a SCAI report

To generate a SCAI Report:
```
scai-report -s <subject artifact filenames> -a <attribute assertion filenames> -o <output filename> [-p <producer filename>]
```

**Note:** The generated SCAI report document is a valid [in-toto Statement].

For a full list of CLI tool options, invoke with the `-h` option.

[in-toto Statement]: https://github.com/in-toto/attestation/blob/main/spec/v1/statement.md
[ResourceDescriptors]: https://github.com/in-toto/attestation/blob/main/spec/v1/resource_descriptor.md
[AttributeAssertions]: https://github.com/in-toto/attestation/blob/main/protos/in_toto_attestation/predicates/scai/v0/scai.proto#L16
[Report]: https://github.com/in-toto/attestation/blob/main/protos/in_toto_attestation/predicates/scai/v0/scai.proto#L28
[SCAI specification]: https://github.com/in-toto/attestation/blob/main/spec/predicates/scai.md
