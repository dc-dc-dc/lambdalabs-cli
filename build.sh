#! /bin/bash

if [ ! -d "out" ]; then
    mkdir out
else
    rm -rf out/*
fi

version="$(git describe --tags --always --abbrev=0 --match='v[0-9]*' 2> /dev/null | sed 's/^.//')"

ldflags=(
    "-X 'main.Version=${version}'"
)

GOOS=linux go build -ldflags="${ldflags[*]}" -o ./out/lambdacli main.go
GOOS=darwin go build -ldflags="${ldflags[*]}" -o ./out/lambdacli.darwin main.go
GOOS=windows go build -ldflags="${ldflags[*]}" -o ./out/lambdacli.exe main.go