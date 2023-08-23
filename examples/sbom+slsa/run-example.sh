#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

VENV_DIR="${VENVDIR:=../../scai-venv}"
EXAMPLE_DIR=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

# -----------------------------------------------------------------
# Run SBOM and SLSA evidence collection example
# -----------------------------------------------------------------

mkdir -p ${EXAMPLE_DIR}/metadata

source ${VENV_DIR}/bin/activate

SBOM_URL="https://github.com/marcelamelara/private-data-objects/suites/15417726142/artifacts/880403395"

PROVENANCE_URL="https://github.com/marcelamelara/private-data-objects/suites/15417726142/artifacts/880403392/pdo_client_wawaka.slsa.intoto.jsonl"

echo GENERATE PDO CLIENT CONTAINER SBOM DESCRIPTOR

scai-gen-resource-desc -n "pdo_client_wawaka.spdx.json" -d -f pdo_client_wawaka.spdx.json -l ${SBOM_URL} -t application/spdx+json --resource-dir ${EXAMPLE_DIR}/metadata -o sbom-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE PDO CLIENT CONTAINER SLSA PROVENANCE DESCRIPTOR

scai-gen-resource-desc -n "build.452e628a.json" -d -f build.452e628a.json -l ${PROVENANCE_URL} -t application/vnd.in-toto.provenance+dsse --resource-dir ${EXAMPLE_DIR}/metadata/attestations -o slsa-desc.json --out-dir ${EXAMPLE_DIR}/metadata

echo GENERATE HAS-SBOM SCAI ATTRIBUTE ASSERTION

scai-attr-assertion -a "HasSBOM" -e ${EXAMPLE_DIR}/metadata/sbom-desc.json -o has-sbom-assertion.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo GENERATE HAS-SLSA SCAI ATTRIBUTE ASSERTION

scai-attr-assertion -a "HasSLSA" -e ${EXAMPLE_DIR}/metadata/slsa-desc.json -o has-slsa-assertion.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print

echo GENERATE SCAI REPORT FOR CONTAINER EVIDENCE COLLECTION

scai-report -s container-img-desc.json --subject-dirs ${EXAMPLE_DIR}/metadata -a has-sbom-assertion.json has-slsa-assertion.json --assertion-dir ${EXAMPLE_DIR}/metadata -o evidence-collection.scai.json --out-dir ${EXAMPLE_DIR}/metadata --pretty-print
