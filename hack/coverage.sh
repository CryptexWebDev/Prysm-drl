#!/usr/bin/env bash

declare -a skip_coverage=("github.com/Dorol-Chain/Prysm-drl/contracts/sharding-manager-contract"
                          "github.com/Dorol-Chain/Prysm-drl/contracts/validator-registration-contract")

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    if [[ ${skip_coverage[*]} =~ $d ]]; then
        continue
    fi
    go test -coverprofile=profile.out -covermode=atomic "$d"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
