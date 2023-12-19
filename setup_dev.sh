#!/usr/bin/env bash

echo "This script assumes that you already installed git, Go 1.19, and Homebrew"

while true; do
    read -p "Do you want to install protobuf compiler and its relevant plugins? " yn
    case $yn in
        [Yy]* ) 
        brew install protobuf
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        break;;
        [Nn]* ) break;;
        * ) echo "Please answer yes or no.";;
    esac
done

while true; do
    read -p "Do you want to install Buf CLI? " yn
    case $yn in
        [Yy]* ) go install github.com/bufbuild/buf/cmd/buf@v1.17.0; break;;
        [Nn]* ) break;;
        * ) echo "Please answer yes or no.";;
    esac
done