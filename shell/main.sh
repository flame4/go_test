#!/usr/bin/env bash


cd ..
CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"
gofmt -w src

# go install leetcode
# go install server
# go install main

# export GOPATH="$OLDGOPATH"
echo "GOPATH=$GOPATH"

