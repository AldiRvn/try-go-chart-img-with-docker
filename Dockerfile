# Stage 1 — Build Go binary
FROM golang:1.25 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o app main.go

# Stage 2 — Runtime pakai headless-shell sebagai dependency
FROM chromedp/headless-shell:latest

WORKDIR /app

# copy binary Go lo
COPY --from=builder /app/app .

# path ke binary chromium headless di image ini
ENV CHROME_PATH=/headless-shell/headless-shell

# JANGAN pakai CMD saja, karena ENTRYPOINT bawaan masih kepake
# Kita override ENTRYPOINT-nya:
ENTRYPOINT ["/app/app"]
