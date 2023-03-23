#! /bin/bash

if [ ! -d "out" ]; then
    mkdir out
else
    rm -rf out/*
fi

GOOS=linux go build -o ./out/lambdacli main.go
GOOS=windows go build -o ./out/lambdacli.exe main.go