# Note

- Don't forget to `go mod tidy; go mod vendor` first
- Docker Run

```bash
mkdir ./gen; docker build -t tws .; docker run \
    -m 1g \
    -p 10822:10822 \
    -v "$(pwd)/gen:/app/gen" \
    tws:latest
```
