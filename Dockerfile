FROM golang:1.18-alpine AS builder

WORKDIR /web-scrawler/

RUN apk update && apk add git upx

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /usr/bin/web-scrawler ./cmd/

RUN upx -9 /usr/bin/web-scrawler

FROM scratch

WORKDIR /web-scrawler/

COPY --from=builder /usr/bin/web-scrawler /usr/bin/web-scrawler
