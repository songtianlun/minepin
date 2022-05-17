FROM golang:1.18.0-alpine3.14 AS builder

# 设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
#   CGO_ENABLED - 纯静态编译

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 编译为二进制可执行文件 app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk upgrade && \
    apk add --no-cache bash git openssh gcc && \
    CGO_ENABLED=1 go install github.com/mattn/go-sqlite3 && \
    make && \
    mv minepin app

# 构建最小镜像
FROM scratch

# 从 builder  中拷贝 /dist/app 到当前目录
COPY --from=builder /build/app /
COPY --from=builder /build/templates /

# 需要运行的命令
ENTRYPOINT ["/app"]