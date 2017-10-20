#!/bin/bash -ex

VERSION=12

rm -f testapp
rm -f testapp.up

CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o testapp .
upx -9 -o testapp.up testapp

docker build -t jtslear/testapp:latest .
docker tag jtslear/testapp:latest jtslear/testapp:cron-1

docker push jtslear/testapp:cron-1
