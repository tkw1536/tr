# tr -- Tiny Redirect

![CI Status](https://github.com/tkw1536/tr/workflows/CI/badge.svg)

This repository contains a tiny http server written in go. 
It implements an HTTP server that answers every incoming request with a redirect.

```TARGET=<TARGET> [ABSOLUTE=1] [PERMANENT=1] [OVERRIDES=/path/to/.json] redirect <bindAddress>```

The address to bind to is specified by the `bindAddress` parameter.

By default, this server redirects all requests to the URL specified by the `TARGET` environment variable.
For each incoming request, the request path is append to it.

For instance, if target is "http://example.com" and the request path is "/index.html", the server will redirect to "http://example.com/index.html".
To disable this behavior and always redirect to the exact `TARGET` set the `ABSOLUTE` environment variable to `1`.

Without additional configuration, all redirect responses return HTTP Status Code 307 (Temporary Redirect).
To instead use Status Code 308 (Permanent Redirect), set the `PERMANENT` environment variable to `1`.

To further override default behavior, the server can also change the behavior for individual request paths.
For this purpose, set the `OVERRIDES` environment variable to a filepath containing a .json file.
This `.json` file is assumed to contain an object mapping request URLS to target URLS.
The request URLS are assumed to have trailing '/'s removed.

## Deployment

This is intended to be used deployed of Docker, and can be found as [a GitHub Package](https://github.com/users/tkw1536/packages/container/package/tr). 
It exposes port 8080 by default.

To start it up run:
```
docker run -e TARGET=http://example.com -p 8080:8080 ghcr.io/tkw1536/tr:latest
```

All `redirect` executable parameters can be passed using standard environment variables.

## Building

This package can be built using standard go tools.
To build the `redirect` executable, install [go 1.15](https://golang.org/dl/) or newer and run:

```bash
go build ./cmd/redirect
```

This will a `redirect` executable in the root directory of the project.

To build the docker image, run:

```bash
docker build -t ghcr.io/tkw1536/tr:latest
```

The docker image is built for `linux/amd64`,`linux/arm64`,`linux/arm/v7` and `linux/arm/v6` architectures by default. 

## Tests

This package contains two tests, go package tests and a docker image test script.
To run the go tests, use:

```bash
go test .
```

To run the docker image tests, run:

```bash
bash test.sh ghcr.io/tkw1536/tr:latest
```

This will build the docker image and then test it.

## Changelog

### Version 1.0.0 (Released [Apr 11 2022](https://github.com/tkw1536/tr/releases/tag/v1.0.0))

- initial release

## License

This project is released into the public domain using the [Unlicense](./LICENSE).
