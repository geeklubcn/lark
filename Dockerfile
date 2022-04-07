FROM golang:1.16 as builder

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/geeklubcn/lark
COPY . $GOPATH/src/github.com/geeklubcn/lark

RUN go build .

EXPOSE 80
ENTRYPOINT ["./lark"]