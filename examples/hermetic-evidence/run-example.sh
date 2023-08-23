#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

VENV_DIR="${VENVDIR:=../../scai-venv}"
EXAMPLE_DIR=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

# -----------------------------------------------------------------
# Run HERMETIC build example
# -----------------------------------------------------------------

mkdir -p ${EXAMPLE_DIR}/metadata

source ${VENV_DIR}/scai-venv/bin/activate

STRACE_LOG_URL="https://github.com/marcelamelara/private-data-objects/suites/15417726142/artifacts/880403398"

echo GENERATE CONTAINER BUILD STRACE LOG DESCRIPTOR

scai-gen-resource-desc -n "strace.log" -d -f strace.log -l ${STRACE_LOG_URL} -t text/plain -a ${EXAMPLE_DIR}/metadata/hermetic-annotation.json --resource-dir ${EXAMPLE_DIR}/metadata -o strace-log-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE IS_HERMETIC_BUILD SCAI ATTRIBUTE ASSERTION

scai-attr-assertion -a "IsHermeticBuild" -t ${EXAMPLE_DIR}/../sbom+slsa/metadata/slsa-desc.json -e ${EXAMPLE_DIR}/metadata/strace-log-desc.json -o is-hermetic-assertion.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo GENERATE SCAI REPORT FOR HERMETIC BUILD REPORT

scai-report -s container-img-desc.json --subject-dirs ${EXAMPLE_DIR}/../sbom+slsa/metadata -a is-hermetic-assertion.json --assertion-dir ${EXAMPLE_DIR}/metadata -o hermetic-build.scai.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print
