printf "in-toto KubeCon + CloudNativeCon NA 2023 demo (verification flow only)\n\n"

# From: https://github.com/marcelamelara/private-data-objects/actions/runs/6740185189/job/18354917647
SLSA_UUID="24296fb24b8ad77ac67df9169ecdd6759b6894daeeafeb95e5398ad34e50418a1e94c6ae9cf7e7d0"

# From: https://github.com/marcelamelara/private-data-objects/actions/runs/6740185189/job/18354917925
SCAI_UUID="24296fb24b8ad77ab2803f68cbf3f73ef6d4c4a0dce5b13e0d86db9f9548fd77de522002a8a8c97c"

printf "Retrieving transparency log entries from Rekor\n\n"
rekor-cli get --uuid $SLSA_UUID > tlog-entries/pdo_client_wawaka.provenance.json
rekor-cli get --uuid $SCAI_UUID > tlog-entries/pdo_client_wawaka.scai.json

printf "Obtaining public keys Rekor log entries\n\n"
scai-gen rekor tlog-entries/pdo_client_wawaka.provenance.json > functionaries/slsa.cert.pem
#scai-gen rekor tlog-entries/pdo_client_wawaka.scai.json > functionaries/scai.cert.pem

printf "Obtaining functionary info\n\n"
in-toto-golang key layout functionaries/slsa.cert.pem > functionaries/slsa.func
#in-toto-golang key layout functionaries/scai.cert.pem > functionaries/scai.func
