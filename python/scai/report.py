#!/usr/bin/env python

# Copyright 2022 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

"""
<Program Name>
  scai.py

<Author>
 Marcela Melara <marcela.melara@intel.com>

<Started>
  Sep 13, 2022

<Copyright>
  See LICENSE for licensing information.

<Purpose>
  Provides a class for SCAI reports, 
  which describe a set of assertions about fine-grained attributes of 
  the subject artifact and its supply chain. SCAI Reports are composable with
  other supply chain metadata formats such as SLSA.
"""

import attr
import securesystemslib.formats as SECSYS_FORMATS
import securesystemslib.exceptions
import securesystemslib.schema as SCHEMA
from in_toto.models.common import Signable
from scai.attribute_assertion import SCAI_ATTRIBUTE_ASSERTION_SCHEMA


SCAI_REPORT_TYPE = "scai/attribute-report/v0.1"

SCAI_REPORT_PRODUCER_SCHEMA = SCHEMA.Object(
    object_name = "SCAI_REPORT_PRODUCER_SCHEMA",
    attributes = SCHEMA.Optional(SCHEMA.ListOf(SCAI_ATTRIBUTE_ASSERTION_SCHEMA))
)


@attr.s(repr=False, init=False)
class Report(Signable):
    """A list of SCAI AttributeAssertions about fine-grained attributes of an artifact or 
    a producer (i.e. its compute environment or platform) in a step of the supply chain.

    Fields:
        subjectAttributes: A list of dictionaries of SCAI Attribute Assertions *asserted* about a
                           subject artifact by the step, i.e::

            [{
              "attribute": "<attribute keyword>",

              "target": A SCAI Object Reference dictionary that describes a target
                  artifact or metadata object to which the asserted attribute applies.

              "conditions": An opaque dictionary that describes specific conditions under 
                  which the associated attribute arises. Acceptable conditions formats are up 
                  to the producer and consumer.

              "evidence": An opaque dictionary that describes evidence metadata, or references
                  another metadata object (via a SCAI Object Reference), for the asserted attribute.
                  It should have at least the "type" (str) entry. 
            },
            ... ]
        producer:
             "attributes": A list of dictionaries of SCAI Attribute Assertions the step.
             "operation":
    """
    _type = attr.ib()
    subjectAttributes = attr.ib()
    producer = attr.ib()

    def __init__(self, **kwargs):
        super(Report, self).__init__()

        self._type = SCAI_REPORT_TYPE
        self.subjectAttributes = kwargs.get("subjectAttributes", [])
        self.producer = kwargs.get("producer", None)
        
        self.validate()

    @property
    def type_(self):
        """The string "scai/report/v0.1" to indentify the SCAI Report type."""
        # NOTE: We expose the type_ attribute in the API documentation instead of
        # _type to protect it against modification.
        # NOTE: Trailing underscore is used by convention (pep8) to avoid conflict
        # with Python's type keyword.
        return self._type

    @staticmethod
    def read(data):
        """Creates a Report object from its dictionary representation.
        
        Arguments:
        data: A dictionary with SCAI Report metadata fields.

        Raises:
        securesystemslib.exceptions.FormatError: Passed data is invalid.

        Returns:
        The created Report object.

        """
        return Report(**data)

    def to_dict(self):
        """Creates a dictionary representation from the Report object.
         Returns:
         The created dictionary with non-null report metadata fields.
        """
        return attr.asdict(self, filter=lambda attr, value: attr.name != '_type' and value != None)

    def _validate_type(self):
        """Private method to check that `_type` is set to SCAI_REPORT_TYPE."""
        if self._type != SCAI_REPORT_TYPE:
            raise securesystemslib.exceptions.FormatError(
                "Invalid Report: field `_type` must be set to 'scai/attribute-report/v0.1', got: {}"
                .format(self._type))
  
    def _validate_subjectAttributes(self):
        """Private method to check that `subjectAttributes` is a `list`."""
        if not isinstance(self.subjectAttributes, list):
            raise securesystemslib.exceptions.FormatError(
                "Invalid Report: field `subjectAttributes` must be of type list, got: {}"
                .format(type(self.subjectAttributes)))
        
        for assertion in list(self.subjectAttributes):
            SCAI_ATTRIBUTE_ASSERTION_SCHEMA.check_match(assertion)

    def _validate_producer(self):
        """Private method to check that `producer` is a `dict`."""
        if self.producer:
            if not isinstance(self.producer, dict):
                raise securesystemslib.exceptions.FormatError(
                    "Invalid Report: field `producer` must be of type dict, got: {}"
                    .format(type(self.producer)))
        
            SCAI_REPORT_PRODUCER_SCHEMA.check_match(self.producer)
            
