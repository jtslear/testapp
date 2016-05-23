#!/bin/bash -ex

VERSION=$(cat VERSION)

rm -f main.up

CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o main .
upx -o main.up main

docker build -t jtslear/testapp:latest .
docker tag jtslear/testapp:latest jtslear/testapp:$VERSION

docker push jtslear/testapp:latest
docker push jtslear/testapp:$VERSION
