# drone-debug

Drone plugin to dump the environment variables and content of filesystem, useful to debug.

## Build

Build the binary with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
```

## Docker

Build the docker image with the following commands:

```
docker build --rm=true -t bclermont/drone-debug .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image.
