#!/usr/bin/env python
# Copyright 2022 Intel Corporation

import json
import os
import errno

# -----------------------------------------------------------------
# -----------------------------------------------------------------
def find_file_in_path(filename, search_path='.') :
    """general utility to search for a file name in a path
    :param str filename: name of the file to locate, absolute path ignores search_path
    :param list(str) search_path: list of directores where the files may be located
    """

    # os.path.abspath only works for full paths, not relative paths
    # this check should catch './abc'
    if os.path.split(filename)[0] :
        if os.path.isfile(filename) :
            return filename
        raise FileNotFoundError(errno.ENOENT, "file does not exist", filename)

    if search_path:
        full_filename = os.path.join(search_path, filename)
        if os.path.isfile(full_filename) :
            return full_filename

    raise FileNotFoundError(errno.ENOENT, "unable to locate file in search path", filename)

def load_json_file(filename, search_path):
    full_file = find_file_in_path(filename, search_path)
    with open(full_file, "r") as rfile :
        contents = rfile.read()
    contents = contents.rstrip('\0')
    return json.loads(contents)

def load_text_file(filename, search_path):
    full_file = find_file_in_path(filename, search_path)
    with open(full_file, "r") as rfile :
        contents = rfile.read()
    contents = contents.rstrip('\0')
    return contents

# read private .pem key file
def parse_pem_file(key_file, search_path):
    full_file = find_file_in_path(key_file, search_path)
    with open(full_file, 'r') as k:
        key = k.read()
    assert key.startswith('-----BEGIN EC PRIVATE KEY-----\n') and key.endswith('\n-----END EC PRIVATE KEY-----\n'), "Malformed .pem key"

    return key

# read public .pem key file
def parse_public_pem_file(key_file, search_path):
    full_file = find_file_in_path(key_file, search_path)
    with open(full_file, 'r') as k:
        key = k.read()
    assert key.startswith('-----BEGIN PUBLIC KEY-----\n') and key.endswith('\n-----END PUBLIC KEY-----\n'), "Malformed .pem key"

    return key