# build stage
# FROM golang:alpine AS build-env
FROM plugins/base:multiarch
ADD . release/helloword /bin/
# RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app


ENTRYPOINT ["/bin/helloword"]