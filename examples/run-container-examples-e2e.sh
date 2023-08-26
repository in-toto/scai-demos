#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

echo RUN SCAI ATTESTATION GENERATION FOR PDO CLIENT CONTAINER EXAMPLE PIPELINE

./examples/sbom+slsa/run-go-example.sh
./examples/hermetic-evidence/run-go-example.sh

echo CHECK PDO CLIENT CONTAINER IN-TOTO LAYOUT

scai-gen check layout -l layouts/pdo_client_wawaka.yml examples/sbom+slsa/metadata/attestations/build.452e628a.json examples/sbom+slsa/metadata/attestations/evidence-collection.1f575092.json

echo CHECK PDO CLIENT CONTAINER HERMETIC BUILD ASSERTION

scai-gen check evidence -p policies/hermetic-evidence.yml -e examples/hermetic-evidence/metadata/ examples/hermetic-evidence/metadata/attestations/build.1f575092.json

echo CHECK PDO CLIENT CONTAINER HAS-SLSA ASSERTION

scai-gen check evidence -p policies/has-slsa.yml -e examples/sbom+slsa/metadata examples/sbom+slsa/metadata/attestations/evidence-collection.1f575092.json

echo NEGATIVE TEST: CHECK BAD PDO CLIENT CONTAINER SCAI ASSERTION

scai-gen check evidence -p policies/hermetic-evidence-fail.yml -e examples/hermetic-evidence/metadata/ examples/hermetic-evidence/metadata/attestations/bad-build.1f575092.json
