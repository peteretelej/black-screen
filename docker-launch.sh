#!/bin/bash
set -e

APP=blackscreen

go vet .
go build -o $APP main.go

docker stop $APP||true
docker rm $APP ||true

docker run -d --name $APP --restart always \
	--net host \
	--memory 50m \
	-v /etc/ssl:/etc/ssl:ro \
	-v /etc/localtime:/etc/localtime:ro \
	-v ${PWD}/${APP}:/${APP} \
	debian:jessie /${APP} -listen ":9521"

docker logs -f $APP


