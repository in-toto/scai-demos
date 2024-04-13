#!/usr/bin/env python
# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

"""
<Program Name>
  scai-gen-resource-desc

<Author>
 Marcela Melara <marcela.melara@intel.com>

<Copyright>
  See LICENSE for licensing information.

<Purpose>
  Command-line interface for generating ITE-9 Resource Descriptors.
"""

import argparse
import json
import os

from scai_generator.utility import load_json_file
from in_toto_attestation.v1.resource_descriptor import ResourceDescriptor

from securesystemslib.hash import digest_filename
import google.protobuf.json_format as pb_json

def Main():
    parser = argparse.ArgumentParser(allow_abbrev=False)

    parser.add_argument('-n', '--name', help='The name', type=str, default='')
    parser.add_argument('-u', '--uri', help='The URI', type=str, default='')
    parser.add_argument('-d', '--digest', help='Flag to generate the SHA256 hash of the resource', action='store_true')
    parser.add_argument('-c', '--content', help='Flag to include the content of the resource', action='store_true')
    parser.add_argument('-l', '--download-location', help='The download location of the resource (if different from the URI)', type=str, default='')
    parser.add_argument('-t', '--media-type', help='The media type of the resource', type=str, default='')
    parser.add_argument('-a', '--annotations', help='Filename of JSON-encoded annotations for the resource descriptor', type=str)
    parser.add_argument('-f', '--resource-file', help='Filename of the resource to be described', type=str)
    parser.add_argument('-o', '--outfile', help='Filename to write out this assertion object', type=str, required=True)
    parser.add_argument('--resource-dir', help='Directory for searching resource files', type=str, default='.')
    parser.add_argument('--annotation-dir', help='Directory for searching annotations files', type=str, default='.')
    parser.add_argument('--out-dir', help='Directory for storing generated files', type=str, default='.')
    parser.add_argument('--pretty-print', help='Flag to pretty-print all json before storing', action='store_true')

    options = parser.parse_args()

    # Fail if the resource descriptor doesn't at least have a name, URI, or digest algorithm specified
    if not options.name and not options.uri and not options.digest:
        print('Error: Need at least one of name, URI or digest for a valid in-toto resource descriptor')
        exit(-1)

    # Generate the digest, if requested
    resource_file_path = None
    if options.resource_file:
        resource_file_path = os.path.join(options.resource_dir, options.resource_file)
    
    resource_digest_set = {}
    if options.digest and resource_file_path:
        # we're ok with using the default hash algorithm (sha256)
        hash_obj = digest_filename(resource_file_path)
        # convert the hash object into a digest set
        resource_digest_set.update({hash_obj.name: hash_obj.hexdigest()})

    resource_bytes = bytes()
    if options.content and resource_file_path:
        with open(resource_file_path, 'rb') as f:
            resource_bytes = f.read()

    annotations_dict = None
    if options.annotations:
        annotations_dict = load_json_file(options.annotations, search_path=options.annotation_dir)

    rd = ResourceDescriptor(name=options.name, uri=options.uri, digest=resource_digest_set, content=resource_bytes, download_location=options.download_location, media_type=options.media_type, annotations=annotations_dict)

    # validate the resource descriptor format
    try:
        rd.validate() 
    except ValueError as e:
        sys.exit(e)
    
    # Write out the resource descriptor file
    outfile = options.outfile
    if not outfile.endswith('.json'):
        outfile += '.json'

    indent = 0
    if options.pretty_print:
        indent = 4

    rd_file = os.path.join(options.out_dir, outfile)
    with open(rd_file, 'w+') as f :
        f.write(pb_json.MessageToJson(rd.pb, indent=indent))

    print('Wrote resource descriptor to %s' % rd_file)

if __name__ == "__main__":
    Main()
