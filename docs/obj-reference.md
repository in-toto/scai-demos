# Object Reference

## Motivation

An Object Reference is designed to be a size-efficient representation of any object, artifact or metadata,
that may be included in any SW supply chain metadata. The Object Reference must allow both humans and automated
verifier programs to easily parse, identify and locate the referenced objects.

For more details, see Section 3 of the [SCAI v0.1 specification]().

## Schema

```
{
    "name": "<NAME>",
    "digest": { "<ALGORITHM>": "VALUE", "...": "..." },
    "locationURI": "<URI>",
    "objectType": "<TYPE>"
}

```

`name` _string, required_

> Human-readable identifier to distinguish the referenced object.
>
> The semantics are up to the producer and consumer. Because consumers may evaluate
> the name against a policy, the name SHOULD be stable between Attestations.
>
> Use: Object lookups.

`digest` _object ([DigestSet](https://github.com/in-toto/Attestation/blob/main/spec/field_types.md#DigestSet)), required_

> Collection of one or more cryptographic digests of the referenced object.
>
> Two DigestSets are considered matching if ANY of the fields match. The
> producer and consumer must agree on acceptable algorithms. If there are no
> overlapping algorithms, the object is considered not matching.
>
> Use: Integrity checks and policy evaluation.

`locationURI` _string ([ResourceURI](https://github.com/in-toto/Attestation/blob/main/spec/field_types.md#ResourceURI)), optional_

> URI to the location of the referenced object.
>
> Acceptable locations (web server, local, git etc.) are up to the producer
> and consumer. To enable a consumer to automatically validate the
> referenced object, the locationURI SHOULD resolve to the object.
>
> Use: Locating and downloading the object matching the `digest`.

`objectType` _string, optional_

> Indicates the type of referenced object.
>
> Acceptable object type formats are up to the producer
> and consumer. Typically, the objectType for an artifact will be its file type.
> The objectType for a metadata object will commonly be a
> data format or schema identifier.
>
> Use: Provide hint about object type, enable type-specific validation.