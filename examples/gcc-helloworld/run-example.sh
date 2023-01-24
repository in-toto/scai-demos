#!/bin/bash

# Copyright 2022 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
CLI_DIR=${SCAI_DIR}/cli
EXAMPLE_DIR=${SCAI_DIR}/examples/gcc-helloworld

# -----------------------------------------------------------------
# Run gcc hello-world example
# -----------------------------------------------------------------

mkdir -p ${EXAMPLE_DIR}/metadata

source ${CLI_DIR}/scai-venv/bin/activate

echo GENERATE STACK PROTECTION SCAI ATTRIBUTE ASSERTION

${CLI_DIR}/scai-attr-assertion -a "WITH_STACK_PROTECTION" -c '{"flags": "-fstack-protector*"}' -o stack-protection-assertion.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo GENERATE SCAI REPORT FOR GCC COMPILATION

GCC_CMD="gcc -fstack-protector -o ${EXAMPLE_DIR}/hello-world ${EXAMPLE_DIR}/hello-world.c"

${CLI_DIR}/scai-report -i hello-world.c -o hello-world --artifact-dir ${EXAMPLE_DIR} -a stack-protection-assertion.json --metadata-dir ${EXAMPLE_DIR}/metadata --pretty-print -c "${GCC_CMD}"
