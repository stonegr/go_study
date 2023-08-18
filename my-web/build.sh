#!/bin/bash

mkdir "Releases"

# 【darwin/amd64】
# echo "start build darwin/amd64 ..."
# CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o ./Releases/m3u8-darwin-amd64 m3u8-downloader.go

# 【linux/amd64】
echo "start build linux/amd64 ..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./Releases/myweb-amd64 ./main

# 【windows/amd64】
echo "start build windows/amd64 ..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./Releases/myweb-amd64.exe ./main

echo "Congratulations,all build success!!!"
