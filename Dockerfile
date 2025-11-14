# Stage 1 — Build Go binary
FROM golang:1.25 AS builder

WORKDIR /app

# Copy semua file termasuk vendor
COPY . .

# Build dengan vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o app main.go



# Stage 2 — Runtime dengan Chromium
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    chromium \
    chromium-common \
    fonts-liberation \
    fonts-noto-color-emoji \
    ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/app /app/app

ENV CHROME_PATH=/usr/bin/chromium

CMD ["/app/app"]
