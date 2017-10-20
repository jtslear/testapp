#!/bin/bash -ex

VERSION=14

rm -f testapp
rm -f testapp.up

CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o testapp .
upx -9 -o testapp.up testapp

sha=$(git rev-parse HEAD)

docker build -t jtslear/testapp:latest .
docker tag jtslear/testapp:latest jtslear/testapp:$VERSION
docker tag jtslear/testapp:latest jtslear/testapp:$sha

docker push jtslear/testapp:latest
docker push jtslear/testapp:$VERSION
docker push jtslear/testapp:$sha
