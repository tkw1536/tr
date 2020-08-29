#!/bin/bash
set -e 

# Build the docker image
IMAGE_NAME="$1"
docker build -t "$IMAGE_NAME" .

STATUS="PASS"

function assert_redirect() {
    GOT=`curl -w "%{url_effective}\n" -I -L -s -S "$1" -o /dev/null`;
    EXPECTED="$2";
    
    if [ "$GOT" == "$EXPECTED" ]; then
        echo -e "\033[0;32m[PASS]\033[0m $1 => $GOT";
    else
        echo -e "\033[0;31m[FAIL]\033[0m $1 => $GOT (expected $EXPECTED)"; 
        STATUS="FAIL"
    fi;
}

docker build -t tkw1536/tr .
docker run -d --rm --name=smoke -p 8080:8080 -e TARGET=https://example.com/subdirectory tkw1536/tr > /dev/null
sleep 10

assert_redirect "http://localhost:8080/" "https://example.com/subdirectory/" 
assert_redirect "http://localhost:8080/path/" "https://example.com/subdirectory/path/"

docker stop smoke > /dev/null

[ "$STATUS" == "PASS" ]