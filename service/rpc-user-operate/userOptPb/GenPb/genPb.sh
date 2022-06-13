goctl rpc protoc UserOptService.proto --go_out=../../ --go-grpc_out=../../ --zrpc_out=../../ --style=goZero --home=../../../../tpl
# 当有多个proto文件时 需要先生成其他的文件 再生成主要的pb文件