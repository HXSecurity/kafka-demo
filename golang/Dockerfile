FROM golang:alpine as builder

WORKDIR /go/src/github.com/HXSecurity/kafka-demo/golang
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && cd cmd/producer/ && go build -gcflags "all=-N -l"

FROM alpine:latest
WORKDIR /dongtai
COPY --from=0 /go/src/github.com/HXSecurity/kafka-demo/golang/cmd/producer ./
EXPOSE 8811
ENTRYPOINT ./producer
