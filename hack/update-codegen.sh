#!/bin/bash

# Copyright 2017 caicloud authors. All rights reserved.

# The script auto-generates kubernetes clients, listers, informers

set -e

ROOT=$(dirname "${BASH_SOURCE}")/..
# Add your packages here
PKGS=("kubeflow/v1alpha1")

CLIENT_PATH=github.com/caicloud/kubeflow-clientset
CLIENT_APIS=$CLIENT_PATH/apis

for path in "${PKGS[@]}"
do
	ALL_PKGS="$CLIENT_APIS/$path "$ALL_PKGS
done

function join {
	local IFS="$1"
   	shift
   	result="$*"
}

join "," ${PKGS[@]}
PKGS=$result

join "," $ALL_PKGS
FULL_PKGS=$result

echo "PKGS: $PKGS"
echo "FULL PKGS: $FULL_PKGS"

cd ${ROOT}

go run cmd/client-gen/main.go \
  -n clientset \
  --clientset-path $CLIENT_PATH \
  --input-base $CLIENT_APIS \
  --input $PKGS

echo "Generated clients"

go run cmd/lister-gen/main.go \
  -o $CLIENT_PATH/listers \
  -i $FULL_PKGS

echo "Generated listers"

go run cmd/informer-gen/main.go

echo "Generated informers"

cd - >/dev/null
