#!/bin/sh

~/go/bin/reflex -r '\.go' -s -- sh -c "APP_PORT=8080 POSTGRES_USER=postgres POSTGRES_PASSWORD=pass SECRET_TOKEN=\"Not so secret token\" go run src/*.go"