#!/usr/bin/env python
# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

"""
<Program Name>
  scai-attr-assertion

<Author>
 Marcela Melara <marcela.melara@intel.com>

<Copyright>
  See LICENSE for licensing information.

<Purpose>
  Command-line interface for generating ITE-9 SCAI Attribute Assertions.
"""

import argparse
import json
import os
import sys

from scai_generator.utility import load_json_file
from in_toto_attestation.predicates.scai.v0.scai import AttributeAssertion
import in_toto_attestation.v1.resource_descriptor_pb2 as rdpb
from in_toto_attestation.v1.resource_descriptor import ResourceDescriptor

import google.protobuf.json_format as pb_json

from securesystemslib.util import get_file_hashes

def Main():
    parser = argparse.ArgumentParser(allow_abbrev=False)

    parser.add_argument('-a', '--attribute', help='Attribute keyword', type=str, required=True)
    parser.add_argument('-c', '--conditions', help='Filename of json-encoded conditions object', type=str)
    parser.add_argument('-e', '--evidence', help='Filename of json-encoded evidence resource descriptor', type=str)
    parser.add_argument('-t', '--target', help='Filename of json-encoded target resource descriptor', type=str)
    parser.add_argument('-o', '--outfile', help='Filename to write out this assertion object', type=str, required=True)
    parser.add_argument('--target-dir', help='Directory for searching target files', type=str)
    parser.add_argument('--evidence-dir', help='Directory for searching evidence files', type=str)
    parser.add_argument('--conditions-dir', help='Directory for searching conditions files', type=str)
    parser.add_argument('--out-dir', help='Directory for storing generated files', type=str)
    parser.add_argument('--pretty-print', help='Flag to pretty-print all json before storing', action='store_true')

    options = parser.parse_args()

    # Create target reference
    target_rd = None
    if options.target:
        target_dict = load_json_file(options.target, search_path=options.target_dir)
        target_rd = pb_json.ParseDict(target_dict, rdpb.ResourceDescriptor())

    # Read conditions
    conditions_dict = None
    if options.conditions:
        conditions_dict = load_json_file(options.conditions, search_path=options.conditions_dir)

    # Read evidence
    evidence_rd = None
    if options.evidence:
        evidence_dict = load_json_file(options.evidence, search_path=options.evidence_dir)
        evidence_rd = pb_json.ParseDict(evidence_dict, rdpb.ResourceDescriptor())

    assertion = AttributeAssertion(options.attribute, target=target_rd, conditions=conditions_dict, evidence=evidence_rd)

    # validate the assertion format, including resource descriptors
    try:
        assertion.validate() 
    except ValueError as e:
        sys.exit(e)

    # Write out the assertions file
    out_dir = '.'
    if options.out_dir:
        out_dir = options.out_dir

    outfile = options.outfile
    if not outfile.endswith('.json'):
        outfile += '.json'

    indent = 0
    if options.pretty_print:
        indent = 4

    assertion_file = os.path.join(out_dir, outfile)
    with open(assertion_file, 'w+') as afile :
        afile.write(pb_json.MessageToJson(assertion.pb, indent=indent))

    print('Wrote attribute assertion to %s' % assertion_file)

if __name__ == "__main__":
    Main()
