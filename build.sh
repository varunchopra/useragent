rm -rf user-agents.json.gz user-agents.json

wget https://github.com/intoli/user-agents/raw/master/src/user-agents.json.gz

gunzip user-agents.json.gz

export CGO_ENABLED=0

# https://github.com/gin-gonic/gin#build-without-msgpack-rendering-feature
go build -tags=nomsgpack .

sudo docker build -t varunchopra/useragent .

sudo docker push varunchopra/useragent
