# STAGE 1: BUILD
FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ ./cmd/
COPY internal/ ./internal/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build --ldflags="-w -s" -o ./bin/godoc-api ./cmd/api/main.go

# STAGE 2: RELEASE
FROM gcr.io/distroless/static-debian12:nonroot-amd64
WORKDIR /app
COPY --from=builder /app/bin/godoc-api ./godoc-api
COPY ./.env ./.env
EXPOSE 8080
HEALTHCHECK --interval=1m --timeout=5s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/v1/alive || exit 1
ENTRYPOINT ["./godoc-api"]
