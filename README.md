# tr -- Tiny Redirect

This repository contains a tiny http server written in go. 
All it can do is redirect unconditionally to a url given as an environment variable. 

This is intended to be used inside of Docker, and can be found as on DockerHub as [tkw01536/tr](https://hub.docker.com/r/tkw01536/tr/) as an automated build. 
To start it up run:

```
docker run -e TARGET=http://example.com -p 8080:80 tkw01536/tr
```

The code is licensed under the Unlicense, hence in the public domain. 