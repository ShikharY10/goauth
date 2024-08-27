# syntax=docker/dockerfile:1

# Build
FROM golang:1.18 AS build               

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o houseware-api ./cmd/main.go

# Deploy
FROM ubuntu

WORKDIR /
COPY --from=build app/houseware-api /houseware-api
COPY --from=build app/.env /.env

EXPOSE 10222

ENTRYPOINT [ "./houseware-api" ]