#!/bin/bash
if [ ! -d ./build ]; then
    mkdir build/
fi
go build -o build/app
