#!/bin/sh
docker run \
    --name nats \
    --network gomicro \
    --rm -p 4222:4222 -p 8222:8222 -d \
    nats