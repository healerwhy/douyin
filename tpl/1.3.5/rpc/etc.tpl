Name: {{.serviceName}}.rpc
ListenOn: 0.0.0.0:9999
# Mode 使用 dev可以通过 grpcui -plaintext localhost:9999 进行调试
Mode: dev
#注册中心
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: {{.serviceName}}.rpc
#如果api不通过etcd，而是直接连接rpc服务
#注意是rpc客户端不是rpc服务端需要配置 Endpoints:
#  - 127.0.0.1:9999
Log:
  Encoding: plain

DB:
  DataSource: username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheConf:
  - Host: localhost:6379
    Pass: Pass
# ----- 一键生成proto文件 sql2pb-----
#sql2pb -go_package ./pb -package pb -host localhost -user user -password password -port 3306 \
#        -schema tablename -service_name xxx  > xxx.proto
# --- docker ----
# goctl docker -go user.go

# ---- model ----
# modeldir=..
# # 数据库配置
# host=localhost
# port=3306
# dbname=douyin
# tablename=user
# username=username
# passwd=passwd
# echo "开始连接库：$dbname 的表：$tablename"
# goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tablename}"  -dir="${modeldir}"
