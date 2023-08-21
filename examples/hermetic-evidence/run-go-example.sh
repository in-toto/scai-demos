#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
EXAMPLE_DIR=${SCAI_DIR}/examples/hermetic-evidence

# -----------------------------------------------------------------
# Run HERMETIC build example
# -----------------------------------------------------------------

OUTDIR=${EXAMPLE_DIR}/metadata/go
mkdir -p ${OUTDIR}

STRACE_LOG_URL="https://github.com/marcelamelara/private-data-objects/suites/15001861846/artifacts/856214824"

PROVENANCE_URL="https://github.com/marcelamelara/private-data-objects/suites/15001861846/artifacts/856214822"

echo GENERATE CONTAINER BUILD STRACE LOG DESCRIPTOR

scai-gen rd file -n "strace.log" -l ${STRACE_LOG_URL} -t text/plain -a ${EXAMPLE_DIR}/metadata/hermetic-annotation.json -o ${OUTDIR}/strace-log-desc.json ${EXAMPLE_DIR}/metadata/strace.log

echo GENERATE PDO CLIENT CONTAINER SLSA PROVENANCE DESCRIPTOR

scai-gen rd file -n "pdo_client_wawaka.provenance.json" -l ${PROVENANCE_URL} -t application/vnd.in-toto.provenance+json -o ${OUTDIR}/slsa-desc.json ${EXAMPLE_DIR}/metadata/pdo_client_wawaka.provenance.json

echo GENERATE IS_HERMETIC_BUILD SCAI ATTRIBUTE ASSERTION

scai-gen assert -t ${OUTDIR}/slsa-desc.json -e ${OUTDIR}/strace-log-desc.json -o ${OUTDIR}/is-hermetic-assertion.json "IsHermeticBuild"

echo GENERATE PDO CLIENT CONTAINER IMAGE DESCRIPTOR

DOCKER_IMG_HASH="03413b4ecb73bba78a71afaf41ac9e921a14b307c9f18fe13344774723a20d82"

scai-gen rd remote -n "pdo_client_wawaka" -g "sha256" -d ${DOCKER_IMG_HASH} -o ${OUTDIR}/container-img-desc.json "example:pdo_client"

echo GENERATE SCAI REPORT FOR HERMETIC BUILD REPORT

scai-gen report -s ${OUTDIR}/container-img-desc.json -o ${OUTDIR}/hermetic-build.scai.json ${OUTDIR}/is-hermetic-assertion.json
