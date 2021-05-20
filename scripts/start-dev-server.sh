#!/bin/sh

~/go/bin/reflex -r '\.go' -s -- sh -c "APP_PATH=127.0.0.1 APP_PORT=8080 SECRET_TOKEN=\"Not so secret token\" go run src/*.go"