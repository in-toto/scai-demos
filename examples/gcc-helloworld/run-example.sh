#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
EXAMPLE_DIR=${SCAI_DIR}/examples/gcc-helloworld

# -----------------------------------------------------------------
# Run gcc hello-world example
# -----------------------------------------------------------------

mkdir -p ${EXAMPLE_DIR}/metadata

source ${SCAI_DIR}/scai-venv/bin/activate

echo RUN GCC

gcc -fstack-protector -o ${EXAMPLE_DIR}/hello-world ${EXAMPLE_DIR}/hello-world.c

echo GENERATE HELLO-WORLD DESCRIPTOR

scai-gen-resource-desc -n hello-world -d -f hello-world --resource-dir ${EXAMPLE_DIR} -o hello-world-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE GCC DESCRIPTOR

GCC_PATH=`which gcc`
GCC_NAME=`gcc --version | head -n 1`

scai-gen-resource-desc -n "${GCC_NAME}" -d -u "file:/${GCC_PATH}" -a cmd-annotation.json --annotation-dir ${EXAMPLE_DIR}/metadata -f ${GCC_PATH} --resource-dir '/' -o gcc-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE STACK PROTECTION SCAI ATTRIBUTE ASSERTION

scai-attr-assertion -a "HAS_STACK_PROTECTION" -c ${EXAMPLE_DIR}/metadata/stack-protector-conditions.json -o stack-protection-assertion.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo GENERATE SCAI REPORT FOR GCC COMPILATION

scai-report -s hello-world-desc.json --subject-dirs ${EXAMPLE_DIR}/metadata -a stack-protection-assertion.json --assertion-dir ${EXAMPLE_DIR}/metadata -p gcc-desc.json --producer-dir ${EXAMPLE_DIR}/metadata -o hello-world.scai.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print
