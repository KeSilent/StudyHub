# 第一阶段：编译二进制文件
FROM golang:alpine AS builder

WORKDIR /go/src/github.com/KeSilent/study-hub/server
COPY go.mod go.sum ./
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# 第二阶段：构建应用镜像
FROM alpine:latest

LABEL MAINTAINER="Jack@kesilent@live.com"

WORKDIR /app
COPY --from=builder /go/src/github.com/KeSilent/study-hub/server/server .
COPY --from=builder /go/src/github.com/KeSilent/study-hub/server/resource ./resource/
COPY --from=builder /go/src/github.com/KeSilent/study-hub/server/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ["./server", "-c", "config.docker.yaml"]
