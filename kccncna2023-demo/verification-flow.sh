printf "in-toto KubeCon + CloudNativeCon NA 2023 demo (verification flow only)\n\n"

printf "DISCLAIMER: This verification flow is only for demo purposes.\n"
printf "A production verification flow includes retrieving and validating the identities/keys of attestation signers, which is not shown in this demo.\n\n"

printf "Verifying ITE-10 Layout\n\n"
attestation-verifier --attestations-directory ./attestations --layout ./policies/layout.yml

printf "\nVerifying SCAI evidence\n\n"
scai-gen check evidence --policy-file ./policies/has-slsa.yml --evidence-dir ./evidence-files ./attestations/evidence-collection.1f575092.json
