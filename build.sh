#!/bin/bash
if [ ! -d ./build ]; then
    mkdir build/
fi
go build -o build/app
if [ -d ./fonts -a ! -d ./build/fonts ]
then
    cp -r ./fonts ./build/fonts
fi
