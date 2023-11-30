#指定基础镜像
FROM daocloud.io/library/ubuntu:20.04

# copy sourcefile to image
WORKDIR /root
COPY . .

# 安装prorobuf编译器及go编译器
RUN apt update && apt install -y protobuf-compiler wget \
    && wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz \
    && rm -rf /usr/local/go && tar -C /usr/local -zxvf go1.21.4.linux-amd64.tar.gz

# 配置go语言环境变量
ENV GOPATH="/root/go"
ENV PATH="/usr/local/go/bin:/root/go/bin:$PATH"

# 安装protoc-gen-go和protoc-gen-go-grpc组件
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 编译并生成可执行程序cache
RUN export PATH=$PATH:$(go env GOPATH)/bin \
    && protoc --go_out=./json_cache ./json_cache/json_cache.proto \
    && protoc --go-grpc_out=./json_cache ./json_cache/json_cache.proto \
    && go build -o cache server.go client.go
