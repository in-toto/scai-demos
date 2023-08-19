#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
EXAMPLE_DIR=${SCAI_DIR}/examples/secure-boot

# -----------------------------------------------------------------
echo "Run TPM2.0 secure boot attestation example"
# -----------------------------------------------------------------

OUTDIR=${EXAMPLE_DIR}/metadata/go
mkdir -p ${OUTDIR}

echo GENERATE LOCALHOST DESCRIPTOR

scai-gen rd remote -n "localhost" -o ${OUTDIR}/localhost-desc.json "http://127.0.0.1"

echo GENERATE MEASURED BOOT REFERENCE STATE DESCRIPTOR

scai-gen rd file -c -t application/json -o ${OUTDIR}/mb-refstate-desc.json ${EXAMPLE_DIR}/metadata/measured_boot_reference_state.json

echo GENERATE MEASURED BOOT REF STATE SCAI ATTRIBUTE ASSERTION

scai-gen assert -t ${OUTDIR}/localhost-desc.json -o ${OUTDIR}/mb-refstate-assertion.json "TPM2.0_REFERENCE_STATE"

echo GENERATE SCAI REPORT FOR TPM REFERENCE VALUE CLAIM

scai-gen report -s ${OUTDIR}/mb-refstate-desc.json -o ${OUTDIR}/tpm-ref-value.scai.json ${OUTDIR}/mb-refstate-assertion.json

echo RUN LOCAL GCC BUILD

gcc -fstack-protector -o ${EXAMPLE_DIR}/../gcc-helloworld/hello-world ${EXAMPLE_DIR}/../gcc-helloworld/hello-world.c

echo GENERATE HELLO-WORLD DESCRIPTOR

scai-gen rd file -n "hello-world" -o ${OUTDIR}/hello-world-desc.json ${EXAMPLE_DIR}/../gcc-helloworld/hello-world

echo GENERATE TPM QUOTE DESCRIPTOR

scai-gen rd file -n "secure boot quote" -t application/json -o ${OUTDIR}/tpm-quote-desc.json ${EXAMPLE_DIR}/metadata/keylime_quote.json

echo GENERATE SECURE BOOT SCAI ATTRIBUTE ASSERTION

scai-gen assert -e ${EXAMPLE_DIR}/metadata/tpm-quote-desc.json -t ${OUTDIR}/localhost-desc.json -o ${OUTDIR}/secure-boot-assertion.json "TAMPER-EVIDENT BUILD PLATFORM"

echo GENERATE SCAI REPORT FOR BUILD WITH SECURE BOOT ATTESTATION

scai-gen report -s ${OUTDIR}/hello-world-desc.json -o ${OUTDIR}/secure-boot.scai.json ${OUTDIR}/secure-boot-assertion.json
