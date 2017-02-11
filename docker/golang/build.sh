#!/bin/bash
set -e
echo "[build.sh:building binary]"
cd $BUILDPATH && go build -o /middleware && rm -rf /tmp/*
echo "[build.sh:launching binary]"
/middleware
