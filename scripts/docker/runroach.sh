#!/bin/sh
docker run -d \
    --name=roach1 \
    -hostname=roach1 \
    --net=gomicro \
    -p 26257:26257 -p 8080:8080  \
    -v "$GOPATH/src/github.com/shyam-unnithan/gomicro/cockroach-data/roach1:/cockroach/cockroach-data"  \
    cockroachdb/cockroach:v20.2.3 start \
    --insecure \
    --join=roach1,roach2,roach3
