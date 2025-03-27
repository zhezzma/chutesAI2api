FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-s -w" -o /app/chutesai2api

FROM alpine:latest

RUN apk add --no-cache \
    ca-certificates \
    tzdata

COPY --from=builder /app/chutesai2api /chutesai2api

EXPOSE 7860
WORKDIR /app/chutesai2api/data
ENTRYPOINT ["/chutesai2api"]
