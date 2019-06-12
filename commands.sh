#!/usr/bin/env bash

set -ex

createGoModule(){
	GO111MODULE=on go mod init
	GO111MODULE=on go build
}

$@