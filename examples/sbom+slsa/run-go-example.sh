#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

EXAMPLE_DIR=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

# -----------------------------------------------------------------
# Run SBOM and SLSA evidence collection example
# -----------------------------------------------------------------

OUTDIR=${EXAMPLE_DIR}/metadata
mkdir -p ${OUTDIR}

SBOM_URL="https://github.com/marcelamelara/private-data-objects/suites/15417726142/artifacts/880403395"

PROVENANCE_URL="https://github.com/marcelamelara/private-data-objects/suites/15417726142/artifacts/880403392/pdo_client_wawaka.slsa.intoto.jsonl"

echo GENERATE PDO CLIENT CONTAINER SBOM DESCRIPTOR

scai-gen rd file -n "pdo_client_wawaka.spdx.json" -l ${SBOM_URL} -t application/spdx+json -o ${OUTDIR}/sbom-desc.json ${EXAMPLE_DIR}/metadata/pdo_client_wawaka.spdx.json

echo GENERATE PDO CLIENT CONTAINER SLSA PROVENANCE DESCRIPTOR

scai-gen rd file -n "build.452e628a.json" -l ${PROVENANCE_URL} -t application/vnd.in-toto+dsse -o ${OUTDIR}/slsa-desc.json ${EXAMPLE_DIR}/metadata/attestations/build.452e628a.json

echo GENERATE HAS-SBOM SCAI ATTRIBUTE ASSERTION

scai-gen assert -e ${OUTDIR}/sbom-desc.json -o ${OUTDIR}/has-sbom-assertion.json "HasSBOM"

echo GENERATE HAS-SLSA SCAI ATTRIBUTE ASSERTION

scai-gen assert -e ${OUTDIR}/slsa-desc.json -o ${OUTDIR}/has-slsa-assertion.json "HasSLSA"

echo GENERATE SCAI REPORT FOR CONTAINER EVIDENCE COLLECTION

scai-gen report -s ${EXAMPLE_DIR}/metadata/container-img-desc.json -o ${OUTDIR}/evidence-collection.scai.json ${OUTDIR}/has-sbom-assertion.json ${OUTDIR}/has-slsa-assertion.json
