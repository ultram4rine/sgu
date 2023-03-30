FROM python:3.9.16 AS init

WORKDIR /init

COPY ./archive.zip ./archive.zip
COPY ./schema.sql ./schema.sql
COPY ./etl.py ./etl.py

RUN ./etl.py

FROM golang:1.19 AS builder

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod download

COPY cmd cmd
COPY internal internal

RUN go build -o eurovision ./cmd/eurovision-parallel/main.go

FROM debian:bullseye

WORKDIR /app

COPY --from=init /init/sqlite.db .
COPY --from=builder /workspace/eurovision .
COPY ./templates ./templates

EXPOSE 3456

ENTRYPOINT ["/app/eurovision"]