#!/bin/sh

~/go/bin/reflex -r '\.go' -s -- sh -c "APP_PATH=127.0.0.1 APP_PORT=4000 go run src/*.go"