#!/bin/bash

set -euo pipefail

command -v swagger >/dev/null 2>&1 || {
  echo >&2 "command 'swagger' not installed. see: https://github.com/go-swagger/go-swagger for installation"
  exit 1
}

scriptPath=$(realpath $0)
scriptDir=$(dirname "$scriptPath")

bggroupDir=$(realpath "$scriptDir/..")

bggroupSpec=$(realpath "$bggroupDir/specs/bggroup.yml")

echo "...generating bggroup server"
swagger generate server -P rest_model_bggroup.Principal -f "$bggroupSpec" -s rest_server_bggroup -t "$bggroupDir" -m "rest_model_bggroup" --exclude-main

echo "...generating bggroup client"
swagger generate client -P rest_model_bggroup.Principal -f "$bggroupSpec" -c rest_client_bggroup -t "$bggroupDir" -m "rest_model_bggroup"

echo "...generating js client"
swagger-codegen generate -i "$bggroupSpec" -l typescript-fetch -o "ui/src/api"