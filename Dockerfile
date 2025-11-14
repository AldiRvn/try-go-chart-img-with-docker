# Stage 1 — Build Go binary
FROM golang:1.25 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o app main.go


# Stage 2 — Runtime dengan Headless Chromium
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    chromium \
    chromium-driver \
    --no-install-recommends \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/app .

ENV CHROME_PATH=/usr/bin/chromium

CMD ["./app"]
