#!/bin/bash -ex

declare -r version=13

declare -r sha=$(git rev-parse HEAD)

docker build -t jtslear/testapp:latest .
docker tag jtslear/testapp:latest jtslear/testapp:$version
docker tag jtslear/testapp:latest jtslear/testapp:$sha

docker push jtslear/testapp:latest
docker push jtslear/testapp:$version
docker push jtslear/testapp:$sha
