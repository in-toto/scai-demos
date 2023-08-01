#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
EXAMPLE_DIR=${SCAI_DIR}/examples/secure-boot

# -----------------------------------------------------------------
echo "Run TPM2.0 secure boot attestation example \n"
# -----------------------------------------------------------------

mkdir -p ${EXAMPLE_DIR}/metadata

source ${SCAI_DIR}/scai-venv/bin/activate

echo GENERATE MEASURED BOOT REFERENCE STATE DESCRIPTOR

scai-gen-resource-desc -n "measured boot reference state" -d -f measured_boot_reference_state.json -t application/json --resource-dir ${EXAMPLE_DIR}/metadata -o mb-refstate-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE TPM QUOTE DESCRIPTOR

scai-gen-resource-desc -d -f keylime_quote.json -t application/json --resource-dir ${EXAMPLE_DIR}/metadata -o tpm-quote-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE MEASURED BOOT REF STATE SCAI ATTRIBUTE ASSERTION

scai-attr-assertion -a "TPM_REFERENCE_STATE" -t ${EXAMPLE_DIR}/metadata/mb-refstate-desc.json -o mb-refstate-assertion.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo GENERATE ATTESTED BOOT SCAI ATTRIBUTE ASSERTION

scai-attr-assertion -a "ATTESTED_BOOT" -e ${EXAMPLE_DIR}/metadata/tpm-quote-desc.json -o attested-boot-assertion.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo GENERATE LOCALHOST DESCRIPTOR

scai-gen-resource-desc -n "localhost" -u "http://127.0.0.1" -o localhost-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE SCAI REPORT FOR SECURE BOOT ATTESTATION

scai-report -s localhost-desc.json --subject-dirs ${EXAMPLE_DIR}/metadata -a mb-refstate-assertion.json attested-boot-assertion.json --assertion-dir ${EXAMPLE_DIR}/metadata -o secure-boot.scai.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print
