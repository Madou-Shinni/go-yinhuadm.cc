# 基础镜像
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

COPY go.mod .
RUN go mod tidy

# 拷贝代码到容器
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod tidy && go env && go build -o server ./cmd/main.go

# 在构建阶段设置执行权限
RUN chmod +x /app/server

# 运行镜像
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
# 复制可执行文件和配置文件
COPY --from=builder /app/server .
COPY --from=builder /app/configs/config.yml ./configs/config.yml

EXPOSE 8888
ENTRYPOINT ["./server", "-f", "./configs/config.yml"]