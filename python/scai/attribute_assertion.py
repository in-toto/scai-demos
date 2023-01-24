#!/usr/bin/env python
# Copyright 2022 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

"""
<Program Name>
  attribute_assertion.py

<Author>
 Marcela Melara <marcela.melara@intel.com>

<Copyright>
  See LICENSE for licensing information.

<Purpose>
  Provides a class for SCAI attribute assertions.
"""

import attr
import securesystemslib.formats as SECSYS_FORMATS
import securesystemslib.exceptions
import securesystemslib.schema as SCHEMA
from in_toto.models.common import Signable
from scai.object_reference import SCAI_OBJ_REF_SCHEMA


SCAI_ATTRIBUTE_ASSERTION_TYPE = "scai/attribute-assertion/v0.1"
    
SCAI_ATTRIBUTE_ASSERTION_SCHEMA = SCHEMA.Object(
    object_name = 'SCAI_ATTRIBUTE_ASSERTION_SCHEMA',
    attribute = SCHEMA.AnyNonemptyString(),
    target = SCHEMA.Optional(SCAI_OBJ_REF_SCHEMA),
    conditions = SCHEMA.Optional(SCHEMA.Object()),
    evidence = SCHEMA.Optional(SCAI_OBJ_REF_SCHEMA))


@attr.s(repr=False, init=False)
class AttributeAssertion(Signable):
    """Attribute Assertions describe functional attributes of a software artifact and its
    supply chain, capable of covering the full software stack of the toolchain that 
    produced the artifact down to the hardware platform. Assertions include information
    about the conditions under which certain functional attributes arise, as well as 
    (authenticated) evidence for the asserted attributes. Together, this information 
    provides a high-integrity description of the functionality of a software artifact and 
    the integrity of its supply chain, which enables a human or program to determine the
    trustworthiness of a given artifact based on specific attributes.

    Fields:
      attribute: A string used to succinctly identify a specific attribute of an artifact.

      target: A SCAI-compliant Object Reference dictionary that describes a target
        artifact or metadata object to which the asserted attribute applies, i.e.::

        {
            "objectType": "<object type descriptor>",
            "name": "<human readable name>",
            "digest": { "<hash algorithm name>": "<hex value of object digest>", ... },
            "locationURI": "<URI to locate the referenced object>"
        }

      conditions: An opaque dictionary that describes specific conditions under 
        which the associated attribute arises. Acceptable conditions formats are up 
        to the producer and consumer.

      evidence: A SCAI-compliant Object Reference dictionary that describes evidence 
        for the asserted attribute. Though not required, providing a hint via the objectType
        filed of the Reference is highly recommended to facilitate verification.

    """
    _type = attr.ib()
    attribute = attr.ib()
    target = attr.ib()
    conditions = attr.ib()
    evidence = attr.ib()
        
    def __init__(self, **kwargs):
        super(AttributeAssertion, self).__init__()

        self._type = SCAI_ATTRIBUTE_ASSERTION_TYPE
        self.attribute = kwargs.get("attribute")
        self.target = kwargs.get("target", None)
        self.conditions = kwargs.get("conditions", None)
        self.evidence = kwargs.get("evidence", None)
        
        self.validate()


    @property
    def type_(self):
        """The string SCAI_ATTRIBUTE_ASSERTION_TYPE to identify the SCAI Attribute Assertion type."""
        # NOTE: We expose the type_ attribute in the API documentation instead of
        # _type to protect it against modification.
        # NOTE: Trailing underscore is used by convention (pep8) to avoid conflict
        # with Python's type keyword.
        return self._type

    @staticmethod
    def read(data):
        """Creates a AttributeAssertion object from its dictionary representation.
         Arguments:
         data: A dictionary with SCAI Attribute Assertion metadata fields.
         Raises:
         securesystemslib.exceptions.FormatError: Passed data is invalid.
         Returns:
         The created AttributeAssertion object.
         """
        return AttributeAssertion(**data)

    def to_dict(self):
        """Creates a dictionary representation from the Attribute Assertion object.
         Returns:
         The created dictionary with non-null attribute assertion metadata fields.
        """
        return attr.asdict(self, filter=lambda attr, value: attr.name != '_type' and value != None)

    def _validate_type(self):
        """Private method to check that `_type` is set to SCAI_ATTRIBUTE_ASSERTION_TYPE."""
        if self._type != SCAI_ATTRIBUTE_ASSERTION_TYPE:
            raise securesystemslib.exceptions.FormatError(
                "Invalid Report: field `_type` must be set to 'scai/attribute-assertion/v0.1', got: {}"
                .format(self._type))

    def _validate_attribute(self):
        """Private method to check that `attribute` is a `str`."""
        if not isinstance(self.attribute, str):
            raise securesystemslib.exceptions.FormatError(
                "Invalid AttributeAssertion: field `attribute` must be of type str, got: {}"
                .format(type(self.attribute)))

    def _validate_target(self):
        """Private method to check that `target` is a `dict`."""
        if self.target:
            if not isinstance(self.target, dict):
                raise securesystemslib.exceptions.FormatError(
                    "Invalid AttributeAssertion: field `target` must be of type dict, got: {}"
                    .format(type(self.target)))
        
            SCAI_OBJ_REF_SCHEMA.check_match(self.target)

    def _validate_conditions(self):
        """Private method to check that `conditions` is a `dict`."""
        if self.conditions:
            if not isinstance(self.conditions, dict):
                raise securesystemslib.exceptions.FormatError(
                    "Invalid AttributeAssertion: field `conditions` must be of type dict, got: {}"
                    .format(type(self.conditions)))

    def _validate_evidence(self):
        """Private method to check that `evidence` is a `dict`."""
        if self.evidence:
            if not isinstance(self.evidence, dict):
                raise securesystemslib.exceptions.FormatError(
                    "Invalid AttributeAssertion: field `evidence` must be of type dict, got: {}"
                    .format(type(self.evidence)))

            SCAI_OBJ_REF_SCHEMA.check_match(self.evidence)


