# tr -- Tiny Redirect

![CI Status](https://github.com/tkw1536/ggman/workflows/CI/badge.svg)

This repository contains a tiny http server written in go. 
All it can do is redirect unconditionally to a url given as an environment variable. 

This is intended to be used inside of Docker, and can be found as [a GitHub Package](https://github.com/tkw1536/tr/packages/377253). 
To start it up run:
```
docker run -e TARGET=http://example.com -p 8080:8080 docker.pkg.github.com/tkw1536/tr/tr:latest
```

For legacy reasons this image is also available on DockerHub as the automated build [tkw01536/tr](https://hub.docker.com/r/tkw01536/tr/). 

The code is licensed under the Unlicense, hence in the public domain. 