# compile stage
FROM golang:1.14.0 AS build-dev

# build delve
RUN go get github.com/go-delve/delve/tree/master/cmd/dlv

ADD . /dockerdev
WORKDIR /dockerdev

RUN go build -gcflags="all=-N -l" -o /server

# Final stage
FROM debian:buster

EXPOSE 8000 40000

WORKDIR /

COPY --from=build-dev /go/bin/dlv /
COPY --from=build-dev /server /

CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/server"]