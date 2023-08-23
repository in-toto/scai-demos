#!/bin/bash

# Copyright 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

echo RUN SCAI ATTESTATION GENERATION FOR PDO CLIENT CONTAINER EXAMPLE PIPELINE

./examples/sbom+slsa/run-go-example.sh
./examples/hermetic-evidence/run-go-example.sh
