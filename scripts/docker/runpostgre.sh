#!/bin/sh
docker run -d \
    --name microdb \
    --network gomicro \
    -e POSTGRES_PASSWORD=verygoodsecret \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v /custom/mount:/var/lib/postgresql/data \
    postgres

