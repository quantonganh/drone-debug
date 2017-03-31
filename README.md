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

## Usage

In your pipeline just add this step to your pipeline

```
pipeline:
  debug:
    image: bclermont/drone-debug
    skip_directories:
      - /proc
      - /dev
```

Output sample:

```
Environment variables
---------------------------
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
DRONE_ARCH=linux/amd64
DRONE_BUILD_CREATED=1490927260
DRONE_BUILD_STATUS=success
HOME=/root

Filesystem
---------------------------
/ isdir:true mode:"drwxr-xr-x" modtime:"2017-03-31 02:27:55.03225343 +0000 UTC" size:4096
/bin isdir:true mode:"drwxr-xr-x" modtime:"2017-03-31 02:25:03 +0000 UTC" size:4096
/etc isdir:true mode:"drwxr-xr-x" modtime:"2017-03-31 02:27:55.016253309 +0000 UTC" size:4096
/etc/apk isdir:true mode:"drwxr-xr-x" modtime:"2017-03-03 11:20:18 +0000 UTC" size:4096
/etc/apk/keys isdir:true mode:"drwxr-xr-x" modtime:"2017-03-03 11:20:17 +0000 UTC" size:4096
/etc/apk/protected_paths.d isdir:true mode:"drwxr-xr-x" modtime:"2017-03-03 11:20:17 +0000 UTC" size:4096
/etc/conf.d isdir:true mode:"drwxr-xr-x" modtime:"2017-03-03 11:20:17 +0000 UTC" size:4096
/etc/crontabs isdir:true mode:"drwxr-xr-x" modtime:"2017-03-03 11:20:17 +0000 UTC" size:4096
```
