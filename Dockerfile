FROM golang:latest AS BUILDER

#设置工作目录
RUN mkdir /app
RUN mkdir /app/iov.tencent.com
RUN mkdir /app/iov.tencent.com/src
RUN mkdir /app/iov.tencent.com/src/configmap-test
RUN mkdir /app/iov.tencent.com/src/configmap-test/config

WORKDIR   /app/iov.tencent.com/src/configmap-test/

ENV GOPATH /app/iov.tencent.com/src/

COPY ./main.go /app/iov.tencent.com/src/configmap-test/
COPY ./go.mod /app/iov.tencent.com/src/configmap-test/
COPY ./go.sum /app/iov.tencent.com/src/configmap-test/
COPY ./config/conf.yml /app/iov.tencent.com/src/configmap-test/config/
COPY ./config/conf-dev.yml /app/iov.tencent.com/src/configmap-test/config/
COPY ./config/config.go /app/iov.tencent.com/src/configmap-test/config/

RUN go build

ENTRYPOINT ["./configmap-test"]