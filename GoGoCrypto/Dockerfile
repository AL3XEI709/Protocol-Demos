# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.20-alpine
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources
# 设置工作目录
WORKDIR /app

# 将本地代码复制到容器中的工作目录
COPY . .

# 构建Go应用程序
RUN go build -o main .

# 暴露Gin框架默认的HTTP端口（通常是8080）
EXPOSE 7842

# 运行应用程序
CMD ["./main"]