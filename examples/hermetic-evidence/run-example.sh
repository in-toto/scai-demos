#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
EXAMPLE_DIR=${SCAI_DIR}/examples/hermetic-evidence

# -----------------------------------------------------------------
# Run HERMETIC build example
# -----------------------------------------------------------------

mkdir -p ${EXAMPLE_DIR}/metadata

source ${SCAI_DIR}/scai-venv/bin/activate

STRACE_LOG_URL="https://github.com/marcelamelara/private-data-objects/suites/15001861846/artifacts/856214824"

PROVENANCE_URL="https://github.com/marcelamelara/private-data-objects/suites/15001861846/artifacts/856214822"

echo GENERATE CONTAINER BUILD STRACE LOG DESCRIPTOR

scai-gen-resource-desc -n "strace.log" -d -f strace.log -l ${STRACE_LOG_URL} -t text/plain -a ${EXAMPLE_DIR}/metadata/hermetic-annotation.json --resource-dir ${EXAMPLE_DIR}/metadata -o strace-log-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE PDO CLIENT CONTAINER SLSA PROVENANCE DESCRIPTOR

scai-gen-resource-desc -n "pdo_client_wawaka.provenance.json" -d -f pdo_client_wawaka.provenance.json -l ${PROVENANCE_URL} -t application/vnd.in-toto.provenance+json --resource-dir ${EXAMPLE_DIR}/metadata -o slsa-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE IS_HERMETIC_BUILD SCAI ATTRIBUTE ASSERTION

scai-attr-assertion -a "IsHermeticBuild" -t ${EXAMPLE_DIR}/metadata/slsa-desc.json -e ${EXAMPLE_DIR}/metadata/strace-log-desc.json -o is-hermetic-assertion.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo GENERATE SCAI REPORT FOR HERMETIC BUILD REPORT

scai-report -s container-img-desc.json --subject-dirs ${EXAMPLE_DIR}/metadata -a is-hermetic-assertion.json --assertion-dir ${EXAMPLE_DIR}/metadata -o hermetic-build.scai.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print
