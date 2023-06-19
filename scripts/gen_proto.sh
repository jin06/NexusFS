#!/usr/bin/env bash

if ! [[ "$0" =~ "scripts/gen_proto.sh" ]]; then
	echo "must be run from repository root"
	exit 255
fi

PRO_ROOT="${PWD}"
PRO_PATH="${PRO_ROOT}/api/"

protoc --go_out=plugins=grpc:. \
  --go_opt=paths=source_relative \
  -I="${PRO_PATH}" \
  ${PRO_PATH}*.proto
