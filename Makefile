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

VENVDIR ?= ./scai-venv
INTOTODIR ?= ../attestation

PY_VERSION=${shell python3 --version | sed 's/Python \(3\.[0-9]\).*/\1/'}
PYTHON_DIR=$(VENVDIR)/lib/python$(PY_VERSION)/site-packages/

$(PYTHON_DIR):
	@echo INSTALL SCAI API
	python3 -m venv $(VENVDIR)
	. $(abspath $(VENVDIR)/bin/activate) && pip install --upgrade pip
	. $(abspath $(VENVDIR)/bin/activate) && pip install --upgrade wheel
	. $(abspath $(VENVDIR)/bin/activate) && pip install --upgrade in-toto
	. $(abspath $(VENVDIR)/bin/activate) && pip install --upgrade ${INTOTODIR}/python
	. $(abspath $(VENVDIR)/bin/activate) && pip install --upgrade ./python

$(VENVDIR):
	@echo CREATE SCAI VENV DIRECTORY $(VENVDIR)
	mkdir -p $(VENVDIR)

py-venv: $(VENVDIR) $(PYTHON_DIR)

go-mod:
	cd ./go && go build && go install

clean:
	@echo REMOVE SCAI VENV AND PYTHON LIB DIRS
	rm -rf $(VENVDIR) __pycache__
	cd ./python; rm -rf build dist *.egg-info

test: py-venv
	@echo TESTING WITH GCC-HELLOWORLD EXAMPLE
	./examples/gcc-helloworld/run-example.sh

.phony: clean test py-venv go-mod
