#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
EXAMPLE_DIR=${SCAI_DIR}/examples/secure-boot

# -----------------------------------------------------------------
echo "Run TPM2.0 secure boot attestation example"
# -----------------------------------------------------------------

mkdir -p ${EXAMPLE_DIR}/metadata

source ${SCAI_DIR}/scai-venv/bin/activate

echo GENERATE LOCALHOST DESCRIPTOR

scai-gen-resource-desc -n "localhost" -u "http://127.0.0.1" -o localhost-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE MEASURED BOOT REFERENCE STATE DESCRIPTOR

scai-gen-resource-desc -d -c -f measured_boot_reference_state.json -t application/json --resource-dir ${EXAMPLE_DIR}/metadata -o mb-refstate-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE MEASURED BOOT REF STATE SCAI ATTRIBUTE ASSERTION

scai-attr-assertion -a "TPM2.0_REFERENCE_STATE" -t localhost-desc.json -o mb-refstate-assertion.json --out-dir ${EXAMPLE_DIR}/metadata --target-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo GENERATE SCAI REPORT FOR TPM REFERENCE VALUE CLAIM

scai-report -s mb-refstate-desc.json --subject-dirs ${EXAMPLE_DIR}/metadata -a mb-refstate-assertion.json --assertion-dir ${EXAMPLE_DIR}/metadata -o tpm-ref-value.scai.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo RUN LOCAL GCC BUILD

gcc -fstack-protector -o ${EXAMPLE_DIR}/../gcc-helloworld/hello-world ${EXAMPLE_DIR}/../gcc-helloworld/hello-world.c

echo GENERATE HELLO-WORLD DESCRIPTOR

scai-gen-resource-desc -n hello-world -d -f hello-world --resource-dir ${EXAMPLE_DIR}/../gcc-helloworld -o hello-world-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE TPM QUOTE DESCRIPTOR

scai-gen-resource-desc -n "secure boot quote" -d -f keylime_quote.json -t application/json --resource-dir ${EXAMPLE_DIR}/metadata -o tpm-quote-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE SECURE BOOT SCAI ATTRIBUTE ASSERTION

scai-attr-assertion -a "TAMPER-EVIDENT BUILD PLATFORM" -e ${EXAMPLE_DIR}/metadata/tpm-quote-desc.json -t ${EXAMPLE_DIR}/metadata/localhost-desc.json -o secure-boot-assertion.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo GENERATE SCAI REPORT FOR BUILD WITH SECURE BOOT ATTESTATION

scai-report -s hello-world-desc.json --subject-dirs ${EXAMPLE_DIR}/metadata -a secure-boot-assertion.json --assertion-dir ${EXAMPLE_DIR}/metadata -o secure-boot.scai.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print
