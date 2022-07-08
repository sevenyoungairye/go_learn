## @sevenyoungairye
## @since 2022/07/08
## docker build -t douban_app:0.1 .
## docker run -itd -p 8999:19999 --name db  douban_app:0.1

# 基于 golang:1.17-alpine
FROM golang:1.17-alpine as builder

# 设置环境变量
ENV GOPATH=/opt/repo
ENV GO111MODULE111=ON
ENV GOPROXY=https://goproxy.io,direct

# 拷贝代码到容器
COPY . /app

# 安装依赖;编译, 将项目放到app目录下
RUN cd /app && go mod download && go build -o /app/doubanNotice

# 第二层
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# COPY 源路径 目标路径 从镜像中 COPY
COPY --from=builder /app /app

# 设置apk仓库
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

# 设置 Asia/Shanghai 时区
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 暴露端口
EXPOSE 19999/tcp

# 启动容器时, 运行项目
CMD ["/app/doubanNotice"]