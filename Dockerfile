FROM golang:alpine
MAINTAINER 1667834841@qq.com
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.io,direct

RUN mkdir /opt/go
WORKDIR /opt/go
COPY . /opt/go/
RUN cd /opt/go
RUN go build -o smzdmPusher
CMD ./smzdmPusher