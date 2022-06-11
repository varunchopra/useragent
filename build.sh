#!/bin/bash -ex
# Dependencies: go, Docker

rm -rf user-agents.json.gz user-agents.json

wget https://github.com/intoli/user-agents/raw/master/src/user-agents.json.gz

gunzip user-agents.json.gz

export CGO_ENABLED=0

# https://github.com/gin-gonic/gin#build-without-msgpack-rendering-feature
go build -tags=nomsgpack .

docker build -t varunchopra/useragent .

docker push varunchopra/useragent
