#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
EXAMPLE_DIR=${SCAI_DIR}/examples/gcc-helloworld

# -----------------------------------------------------------------
# Run gcc hello-world example
# -----------------------------------------------------------------

OUTDIR=${EXAMPLE_DIR}/metadata/go

mkdir -p ${OUTDIR}

echo RUN GCC

gcc -fstack-protector -o ${EXAMPLE_DIR}/hello-world ${EXAMPLE_DIR}/hello-world.c

echo GENERATE HELLO-WORLD DESCRIPTOR

scai-gen rd file -n hello-world -o ${OUTDIR}/hello-world-desc.json ${EXAMPLE_DIR}/hello-world

echo GENERATE GCC DESCRIPTOR

GCC_PATH=`which gcc`
GCC_NAME=`gcc --version | head -n 1`

scai-gen rd file -n "${GCC_NAME}" -o ${OUTDIR}/gcc-desc.json ${GCC_PATH}

echo GENERATE STACK PROTECTION SCAI ATTRIBUTE ASSERTION

scai-gen assert -c ${EXAMPLE_DIR}/metadata/stack-protector-conditions.json -o ${OUTDIR}/stack-protection-assertion.json "HAS_STACK_PROTECTION"

echo GENERATE SCAI REPORT FOR GCC COMPILATION

scai-gen report -s ${OUTDIR}/hello-world-desc.json -p ${OUTDIR}/gcc-desc.json -o ${OUTDIR}/hello-world.scai.json ${OUTDIR}/stack-protection-assertion.json
