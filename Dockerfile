FROM golang:1.16.4-alpine3.13 as builder

WORKDIR /go/src/github.com/kamaal111/metrics-service
RUN go mod download
COPY . /go/src/github.com/kamaal111/metrics-service
COPY . entrypoint.sh
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/metrics-service github.com/kamaal111/metrics-service/src

FROM alpine
COPY --from=builder /go/src/github.com/kamaal111/metrics-service /usr/bin/metrics-service/src
EXPOSE 8080 8080
RUN chmod +x /usr/bin/metrics-service
ENTRYPOINT ["/usr/bin/metrics-service"]