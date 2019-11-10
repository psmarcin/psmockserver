#!/bin/bash
set -o nounset
set -e

name="psmarcin/psmockserver:latest"
docker build ./ -t "${name}"
dockerId="$(docker run -d -p 8080:8080 ${name})"
echo "🚀 Starting mock tests..."
go test ./test/...
echo "✅ Finished mock tests"
echo "🧤  Cleaning up..."
docker stop ${dockerId}
echo "✅ Cleaning finished"
