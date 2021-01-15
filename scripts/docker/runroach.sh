#!/bin/sh
docker run -d \
    --name=roach1 \
    --hostname=roach1 \
    --net=gomicro \
    -p 26257:26257 -p 8080:8080  \
    -v "$GOPATH/src/github.com/shyam-unnithan/go-micro/cockroach-data/roach1:/cockroach/cockroach-data"  \
    cockroachdb/cockroach:v20.2.3 start \
    --insecure \
    --join=roach1,roach2,roach3

docker run -d \
    --name=roach2 \
    --hostname=roach2 \
    --net=gomicro \
    -v "$GOPATH/src/github.com/shyam-unnithan/go-micro/cockroach-data/roach2:/cockroach/cockroach-data"  \
    cockroachdb/cockroach:v20.2.3 start \
    --insecure \
    --join=roach1,roach2,roach3

docker run -d \
    --name=roach3 \
    --hostname=roach3 \
    --net=gomicro \
    -v "$GOPATH/src/github.com/shyam-unnithan/go-micro/cockroach-data/roach3:/cockroach/cockroach-data"  \
    cockroachdb/cockroach:v20.2.3 start \
    --insecure \
    --join=roach1,roach2,roach3

    docker exec -it roach1 ./cockroach init --insecure

