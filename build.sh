#!/bin/bash
# Easyjson 命令行构建处理 chmod +x build.sh

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/chatai.exe

echo build windows success

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/chatai

echo build linux success
