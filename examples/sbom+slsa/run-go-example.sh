#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
EXAMPLE_DIR=${SCAI_DIR}/examples/sbom+slsa

# -----------------------------------------------------------------
# Run SBOM and SLSA evidence collection example
# -----------------------------------------------------------------

OUTDIR=${EXAMPLE_DIR}/metadata/go
mkdir -p ${OUTDIR}

SBOM_URL="https://github.com/marcelamelara/private-data-objects/suites/14743599027/artifacts/838001247"

PROVENANCE_URL="https://github.com/marcelamelara/private-data-objects/suites/14743599027/artifacts/838001246"

echo GENERATE PDO CLIENT CONTAINER SBOM DESCRIPTOR

scai-gen rd file -n "pdo_client SBOM" -l ${SBOM_URL} -t application/spdx+json -o ${OUTDIR}/sbom-desc.json ${EXAMPLE_DIR}/metadata/pdo_client_wawaka.spdx.json

echo GENERATE PDO CLIENT CONTAINER SLSA PROVENANCE DESCRIPTOR

scai-gen rd file -n "pdo_client SLSA" -l ${PROVENANCE_URL} -t application/vnd.in-toto.bundle+jsonl -o ${OUTDIR}/slsa-desc.json ${EXAMPLE_DIR}/metadata/pdo_client_wawaka.slsa.intoto.jsonl

echo GENERATE HAS-SBOM SCAI ATTRIBUTE ASSERTION

scai-gen assert -e ${OUTDIR}/sbom-desc.json -o ${OUTDIR}/has-sbom-assertion.json "HasSBOM"

echo GENERATE HAS-SLSA SCAI ATTRIBUTE ASSERTION

scai-gen assert -e ${OUTDIR}/slsa-desc.json -o ${OUTDIR}/has-slsa-assertion.json "HasSLSA"

echo GENERATE PDO CLIENT CONTAINER IMAGE DESCRIPTOR

DOCKER_IMG_HASH="f2d58a9a85dbc80fcd6c52964c9b24b2dce8f6e6e1cdd65c48e8d109dde7f0e4"

scai-gen rd remote -n "pdo_client_wawaka" -g "sha256" -d ${DOCKER_IMG_HASH} -o ${OUTDIR}/container-img-desc.json "example:pdo_client"

echo GENERATE SCAI REPORT FOR CONTAINER EVIDENCE COLLECTION

scai-gen report -s ${OUTDIR}/container-img-desc.json -o ${OUTDIR}/evidence-collection.scai.json ${OUTDIR}/has-sbom-assertion.json ${OUTDIR}/has-slsa-assertion.json
