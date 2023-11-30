下面两个命令执行前要安装go install google.golang.org/protobuf/cmd/protoc-gen-go@latest和
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        还要通过命令export PATH=$PATH:$(go env GOPATH)/bin来添加go的环境变量（通过这个命令export PATH=$PATH:$GOPATH/bin是不行的）
进入proto所在文件路径，使用 protoc --go_out=. proto文件名  命令可以生成go语言相关的文件,如client.pb.go
进入proto所在文件路径，使用 protoc --go-grpc_out=. proto文件名  命令可以生成rpc相关的文件,如client_grpc.pb.go
Protobuf3 语法:
reserved（保留标识符）:可使用reserved告诉编译器和其他开发者在这个消息中不能使用这些特定的字段号或字段名。这有助于避免潜在的冲突，特别是在多个团队合作或在不同版本之间演变时。
singular：一个格式良好的消息应该有0个或者1个这种字段（但是不能超过1个）。
repeated：在一个格式良好的消息中，这种字段可以重复任意多次（包括0次）。重复的值的顺序会被保留。
定义.proto文件时能够标注一系列的option。Option并不改变整个文件声明的含义，但却能够影响特定环境下处理方式。