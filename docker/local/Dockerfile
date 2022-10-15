FROM golang:1.18-alpine AS builder

WORKDIR /archie

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o archie main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /archie

COPY --from=builder /archie/archie ./
COPY ./docker/local/archie.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478

ENTRYPOINT ["./archie.sh"]
