#!/bin/bash

# Copyright 2022 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
VENV=${SCAI_DIR}/cli/scai-venv/bin/activate

if [ ! -f "$VENV" ]; then
    echo "Need to setup SCAI virtualenv first. Please run `make -C cli` from the SCAI root directory first."
    exit -1
fi

source ${VENV}

# Install atheris and coverage
pip install --upgrade atheris
pip install --upgrade coverage

# Run the JSON input fuzzer
echo RUN THE SCAI FUZZER

python3 -m coverage run scai_fuzz.py -atheris_runs=10000000 -max_len=8192 > fuzz-output.txt 2>&1
