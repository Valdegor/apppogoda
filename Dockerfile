FROM golang:alpine AS builder

WORKDIR /app

COPY main.go .

RUN go mod init pogoda

RUN go build -ldflags="-s -w" -o app

FROM scratch

LABEL org.opencontainers.image.authors="Bohdan Maikut"

COPY --from=builder /app/app /app

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget --spider --quiet http://localhost:8080/ || exit 1

ENTRYPOINT ["/app"]
