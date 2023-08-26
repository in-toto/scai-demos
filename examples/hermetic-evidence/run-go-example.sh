#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

EXAMPLE_DIR=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

# -----------------------------------------------------------------
# Run HERMETIC build example
# -----------------------------------------------------------------

OUTDIR=${EXAMPLE_DIR}/metadata
mkdir -p ${OUTDIR}

STRACE_LOG_URL="https://github.com/marcelamelara/private-data-objects/suites/15417726142/artifacts/880403398"

echo GENERATE CONTAINER BUILD STRACE LOG DESCRIPTOR

scai-gen rd file -n "strace.log" -l ${STRACE_LOG_URL} -t text/plain -a ${EXAMPLE_DIR}/metadata/hermetic-annotation.json -o ${OUTDIR}/strace-log-desc.json ${EXAMPLE_DIR}/metadata/strace.log

echo GENERATE IS_HERMETIC_BUILD SCAI ATTRIBUTE ASSERTION

scai-gen assert -t ${EXAMPLE_DIR}/../sbom+slsa/metadata/slsa-desc.json -e ${OUTDIR}/strace-log-desc.json -o ${OUTDIR}/is-hermetic-assertion.json "IsHermeticBuild"

echo GENERATE SCAI REPORT FOR HERMETIC BUILD REPORT

scai-gen report -s ${EXAMPLE_DIR}/../sbom+slsa/metadata/container-img-desc.json -o ${OUTDIR}/hermetic-build.scai.json ${OUTDIR}/is-hermetic-assertion.json

echo GENERATE NON_HERMETIC_BUILD SCAI ATTRIBUTE ASSERTION

scai-gen assert -t ${EXAMPLE_DIR}/../sbom+slsa/metadata/slsa-desc.json -e ${OUTDIR}/strace-log-desc.json -o ${OUTDIR}/non-hermetic-assertion.json "NonHermeticBuild"

echo GENERATE SCAI REPORT FOR HERMETIC BUILD REPORT

scai-gen report -s ${EXAMPLE_DIR}/../sbom+slsa/metadata/container-img-desc.json -o ${OUTDIR}/non-hermetic-build.scai.json ${OUTDIR}/non-hermetic-assertion.json
