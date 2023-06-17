# Copyright 2023 Intel Corporation
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

DSTDIR=scai-venv

PY_VERSION=${shell python3 --version | sed 's/Python \(3\.[0-9]\).*/\1/'}
PYTHON_DIR=$(DSTDIR)/lib/python$(PY_VERSION)/site-packages/

all : environment

$(PYTHON_DIR) :
	@echo INSTALL SCAI API
	python3 -m venv $(DSTDIR)
	. $(abspath $(DSTDIR)/bin/activate) && pip install --upgrade pip
	. $(abspath $(DSTDIR)/bin/activate) && pip install --upgrade wheel
	. $(abspath $(DSTDIR)/bin/activate) && pip install --upgrade in-toto
	. $(abspath $(DSTDIR)/bin/activate) && pip install --upgrade ../../in-toto.attestation/python
	. $(abspath $(DSTDIR)/bin/activate) && pip install --upgrade ../python

$(DSTDIR) :
	@echo CREATE SCAI VENV DIRECTORY $(DSTDIR)
	mkdir -p $(DSTDIR)

environment: $(DSTDIR) $(PYTHON_DIR)

clean:
	@echo REMOVE SCAI VENV AND PYTHON LIB DIRS
	@rm -rf $(DSTDIR) __pycache__
	@cd ../python; rm -rf build dist *.egg-info

.phony : all
.phony : clean
.phony : environment
