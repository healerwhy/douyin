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
  DataSource: healer:healer000.@tcp(120.79.222.123:3306)/go-zero-stud?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheConf:
- Host: 120.79.222.123:6379
  Pass: healer000.

# --- docker ----
# goctl docker -go user.go

# ---- model ----
# modeldir=..
# # 数据库配置
# host=120.79.222.123
# port=3306
# dbname=go-zero-stud
# tablename1=user
# tablename2=user_data
# username=healer
# passwd=healer000.
# echo "开始连接库：$dbname 的表：$tablename1"
# goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tablename1}"  -dir="${modeldir}"
