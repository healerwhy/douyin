goctl api go -api userOpt.api -dir ../ --style=goZero --home=../../../tpl
# 生成api业务代码 ， 进入"服务/cmd/api/desc"目录下，执行下面命令
# goctl api go -api *.api -dir ../  --style=goZero

# 生成rpc业务代码
# 【注】 需要安装下面3个插件
#       protoc >= 3.13.0 ， 如果没安装请先安装 https://github.com/protocolbuffers/protobuf，下载解压到$GOPATH/bin下即可，前提是$GOPATH/bin已经加入$PATH中
#       protoc-gen-go ，如果没有安装请先安装 go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#       protoc-gen-go-grpc  ，如果没有安装请先安装 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
#
#       如果有要使用grpc-gateway，也请安装如下两个插件 , 没有使用就忽略下面2个插件
#       go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
#       go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
#
# 1）goctl >= 1.3 进入"服务/cmd/rpc-user-info/userInfoPb"目录下，执行下面命令
#    goctl rpc-user-info protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../ --style=goZero
#    去除proto中的json的omitempty
#    mac: sed -i "" 's/,omitempty//g' *.userInfoPb.go
#    linux: sed -i 's/,omitempty//g' *.userInfoPb.go
# 2）goctl < 1.3 进入"服务/cmd"目录下，执行下面命令
#    goctl rpc-user-info proto -src rpc-user-info/userInfoPb/*.proto -dir ./rpc-user-info --style=goZero
#    去除proto中的json的omitempty
#    mac: sed -i "" 's/,omitempty//g'  ./rpc-user-info/userInfoPb/*.userInfoPb.go
#    linux: sed -i 's/,omitempty//g'  ./rpc-user-info/userInfoPb/*.userInfoPb.go

# 创建kafka的topic
# kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 -partitions 1 --topic {topic}
# 查看消费者组情况
# kafka-consumer-groups.sh --bootstrap-server kafka:9092 --describe --group {group}
# 命令行消费
# ./kafka-console-consumer.sh  --bootstrap-server kafka:9092  --topic looklook-log   --from-beginning
# 命令生产
# ./kafka-console-producer.sh --bootstrap-server kafka:9092 --topic second
