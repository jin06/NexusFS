#!/usr/bin/env bash

if ! [[ "$0" =~ "scripts/gen_proto.sh" ]]; then
	echo "must be run from repository root"
	exit 255
fi

PRO_ROOT="${PWD}"
PRO_PATH="${PRO_ROOT}/api/"
PRO_OUT="${PRO_ROOT}/api/"

protoc  \
  --proto_path=${PRO_PATH} \
  --go_out=plugins=grpc:${PRO_OUT} \
  --go_opt=paths=source_relative \
  ${PRO_PATH}*.proto
