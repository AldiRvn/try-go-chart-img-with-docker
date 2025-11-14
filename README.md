# Note

- Don't forget to `go mod tidy; go mod vendor` first
- Docker Run

```bash
rm -drf ./gen;mkdir ./gen; docker build -t tws .; docker run \
    -m 1g \
    -p 10822:10822 \
    -v "$(pwd)/gen:/app/gen" \
    tws:latest
```

- Check Docker Image Size

```bash
docker images | grep tws

WARNING: This output is designed for human readability. For machine-readable output, please use --format.
tws:latest                      4f4384aa556d        327MB             0B   U    
```
