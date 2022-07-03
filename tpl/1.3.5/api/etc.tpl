Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}

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
# tablename1=user
# tablename2=user_data
# username=username
# passwd=passwd
# echo "开始连接库：$dbname 的表：$tablename1"
# goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tablename1}"  -dir="${modeldir}"