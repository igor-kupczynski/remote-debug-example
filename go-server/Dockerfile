# Build
FROM golang:1.14 AS build-env

ADD . /workspace
WORKDIR /workspace

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Run
FROM alpine:3.11

EXPOSE 8080

WORKDIR /
COPY --from=build-env /workspace/app /app

CMD ["/app"]