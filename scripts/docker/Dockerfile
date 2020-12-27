# Multistage docker build file
########################################
# STEP 1 build executable binary
########################################
FROM golang:1.15.6-alpine AS builder


# Create appuser
ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN apk update && apk add --no-cache git make 
