#!/bin/bash
set -o nounset
set -e

docker kill $(docker ps --filter ancestor=psmarcin/psmockserver:latest -q)
