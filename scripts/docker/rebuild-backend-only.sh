#!/bin/sh

cd ../../
go build .
cd scripts/docker
sudo docker compose restart backend