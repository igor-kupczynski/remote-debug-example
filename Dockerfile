# Build
FROM golang:1.14 AS build-env

ADD . /workspace
WORKDIR /workspace

RUN go build -o app

# Run
FROM ubuntu:20.04

EXPOSE 8080

WORKDIR /
COPY --from=build-env /workspace/app /app

CMD ["/app"]