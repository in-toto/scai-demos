#!/usr/bin/env python
# Copyright 2022 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

import sys
import atheris
import json
from collections.abc import Mapping
import securesystemslib.exceptions

with atheris.instrument_imports():
    from scai.report import Report
    from scai.attribute_assertion import AttributeAssertion
    from scai.object_reference import ObjectReference

def TestOneInput(input_bytes):
    fdp = atheris.FuzzedDataProvider(input_bytes)
    data = fdp.ConsumeUnicode(sys.maxsize)
    
    try:
       json_dict = json.loads(data)
    except json.JSONDecodeError:
        return

    if not isinstance(json_dict, Mapping):
        return

    try:
        ObjectReference.read(json_dict)
        AttributeAssertion.read(json_dict)
        Report.read(json_dict)
    except securesystemslib.exceptions.FormatError:
        pass
    
atheris.Setup(sys.argv, TestOneInput)
atheris.Fuzz()