#!/bin/bash

basedir=$(realpath $(dirname $0))

# Generate DB models
cd ${basedir}/pkg/db && oapi-codegen --config models.cfg.yaml ${basedir}/swagger.yaml

# Generate server model types
cd ${basedir}/pkg/web/api && oapi-codegen --config types.cfg.yaml ${basedir}/swagger.yaml

# Generate server endpoints
cd ${basedir}/pkg/web/api && oapi-codegen --config server.cfg.yaml ${basedir}/swagger.yaml

cd ${basedir}
