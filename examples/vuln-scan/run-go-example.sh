#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

SCAI_DIR=~/supply-chain-attribute-integrity
EXAMPLE_DIR=${SCAI_DIR}/examples/vuln-scan

# -----------------------------------------------------------------
echo "Run Snyk scan attestation example"
# -----------------------------------------------------------------

OUTDIR=${EXAMPLE_DIR}/metadata
mkdir -p ${OUTDIR}

SNYK_LOG_LOCATION="https://scans.example.com/results/snyk"

echo GENERATE SNYK LOG DESCRIPTOR

scai-gen rd file -n "snyk-results.txt" -l ${SNYK_LOG_LOCATION} -t text/plain -o ${OUTDIR}/snyk-log-desc.json ${EXAMPLE_DIR}/metadata/snyk-results.txt

echo GENERATE VULN SCAN ATTRIBUTE ASSERTION

scai-gen assert -e ${OUTDIR}/snyk-log-desc.json -o ${OUTDIR}/no-vulns-assertion.json "NO KNOWN VULNERABILITIES"

echo GENERATE GIT REPO DESCRIPTOR

scai-gen rd remote -g sha1 -d caf62c3caf2ab1228fb546f149f26df23c308802 -o ${OUTDIR}/repo-desc.json "https://github.com/IntelLabs/supply-chain-attribute-integrity"

echo GENERATE VULN SCAN ATTRIBUTE REPORT

scai-gen report -s ${OUTDIR}/repo-desc.json -o ${OUTDIR}/no-vuln.scai.json ${OUTDIR}/no-vulns-assertion.json
