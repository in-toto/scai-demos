#!/usr/bin/env python
# Copyright 2022 Intel Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import re

import os
import shutil

# this should only be run with python3
import sys
if sys.version_info[0] < 3:
    print('ERROR: must run with python3')
    sys.exit(1)

from setuptools import setup, find_packages, Extension

setup(name='scai-generator',
      version='0.2',
      description='SCAI metadata generation tools',
      author='Intel Corporation',
      packages=find_packages(),
      entry_points = {
          'console_scripts': [
              'scai-attr-assertion = scai_generator.cli.gen_attr_assertion:Main',
              'scai-gen-resource-desc = scai_generator.cli.gen_resource_desc:Main',
              'scai-report = scai_generator.cli.gen_report:Main'
          ]
      }
)
