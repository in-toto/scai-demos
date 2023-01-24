#!/usr/bin/env python

# Copyright 2022 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

"""
<Program Name>
  object_reference.py

<Author>
  Marcela Melara <marcela.melara@intel.com>

<Started>
  Oct 3, 2022

<Copyright>
  See LICENSE for licensing information.

<Purpose>
  Provides a class for object references according to the SCAI specification.
  A reference is a 4-tuple data structure that describes the name, cryptographic 
  digest set, resolvable location and type of an object, enabling consumers to 
  locate and validate the described object.
"""

import attr
import securesystemslib.formats as SECSYS_FORMATS
import securesystemslib.exceptions
import securesystemslib.schema as SCHEMA
from in_toto.models.common import Signable


SCAI_OBJ_REF_SCHEMA = SCHEMA.Object(
    object_name = 'SCAI_OBJ_REF_SCHEMA',
    name = SECSYS_FORMATS.NAME_SCHEMA,
    digest = SECSYS_FORMATS.HASHDICT_SCHEMA,
    locationUri = SCHEMA.Optional(SECSYS_FORMATS.URL_SCHEMA),
    objectType = SCHEMA.Optional(SCHEMA.AnyString()))
    

@attr.s(repr=False, init=False)
class ObjectReference(Signable):
    """An Object Reference is designed to be a size-efficient representation of any 
    object, artifact or metadata, that may be included in any metadata object. The Object 
    Reference must allow both humans and automated verifier programs to easily parse, 
    identify and locate the referenced objects.

    Attributes:
      name: A string used to succinctly identify a specific name of an artifact.

      digest: A set of cryptographic digests of the object.

      locationUri: Optional URI string pointing to the location of the object.

      objectType: Optional string indicating the type of referenced object.
    """
    name = attr.ib()
    digest = attr.ib()
    locationUri = attr.ib()
    objectType = attr.ib()
    
    
    def __init__(self, **kwargs):
        super(ObjectReference, self).__init__()

        self.name = kwargs.get("name")
        self.digest = kwargs.get("digest", {})
        self.locationUri = kwargs.get("locationUri", None)
        self.objectType = kwargs.get("objectType", None)
        
        self.validate()

    @staticmethod
    def read(data):
        """Creates an Object Reference object from its dictionary representation.
         Arguments:
         data: A dictionary with object reference metadata fields.
         Raises:
         securesystemslib.exceptions.FormatError: Passed data is invalid.
         Returns:
         The created ObjectReference object.
        """
        return ObjectReference(**data)

    def to_dict(self):
        """Creates a dictionary representation from the Object Reference object.
         Returns:
         The created dictionary with non-null object reference metadata fields.
        """
        return attr.asdict(self, filter=lambda attr, value: value != None)
        
    def _validate_name(self):
        """Private method to check that `name` is a `str`."""
        if not isinstance(self.name, str):
            raise securesystemslib.exceptions.FormatError(
                "Invalid ObjectReference: field `name` must be of type str, got: {}"
                .format(type(self.name)))

    def _validate_digest(self):
        """Private method to check that `digest` is a `HASHDICT`."""
        if not isinstance(self.digest, dict):
            raise securesystemslib.exceptions.FormatError(
                "Invalid ObjectReference: field `digest` must be of type dict, got: {}"
                .format(type(self.digest)))
        SECSYS_FORMATS.HASHDICT_SCHEMA.check_match(self.digest)

    def _validate_locationUri(self):
        """Private method to check that `locationUri` is a `URL str`."""
        if self.locationUri:
            if not isinstance(self.locationUri, str):
                raise securesystemslib.exceptions.FormatError(
                    "Invalid ObjectReference: field `locationUri` must be of type str, got: {}"
                    .format(type(self.locationUri)))

            SECSYS_FORMATS.URL_SCHEMA.check_match(self.locationUri)

    def _validate_objectType(self):
        """Private method to check that `objectType` is a `str`."""
        if self.objectType:
            if not isinstance(self.objectType, str):
                raise securesystemslib.exceptions.FormatError(
                    "Invalid ObjectReference: field `objectType` must be of type str, got: {}"
                    .format(type(self.objectType)))
