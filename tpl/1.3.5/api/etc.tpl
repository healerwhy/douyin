Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}

Log:
  Encoding: plain

DB:
  DataSource: healer:healer000.@tcp(120.79.222.123:3306)/go-zero-stud?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheConf:
  - Host: 120.79.222.123:6379
    Pass: healer000.
# ----- 一键生成proto文件 sql2pb-----
#sql2pb -go_package ./pb -package pb -host 120.79.222.123 -user healer -password healer000. -port 3306 \
#        -schema go-zero-stud -service_name GetUserService  > user.proto
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