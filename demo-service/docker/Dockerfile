
# 使用 golang 官方镜像提供 Go 运行环境，并且命名为 buidler 以便后续引用
FROM golang:1.14-alpine as builder

# 启用 Go Modules 并设置 GOPROXY
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache bash supervisor curl

# 设置工作目录
RUN mkdir -p /working
WORKDIR /working
Copy ./demo-service /working

Copy ./demo-service/docker/supervisord.conf /etc/supervisord.conf


# 构建二进制文件，添加来一些额外参数以便可以在 Alpine 中运行它
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main  /working/main.go
  
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
