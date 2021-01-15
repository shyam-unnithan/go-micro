#!/bin/sh
docker run -it --rm --network gomicro postgres psql -h microdb -U postgres
