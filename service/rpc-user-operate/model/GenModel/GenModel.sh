#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package

#表生成的genmodel目录
modeldir=..

# 数据库配置
host=120.79.222.123
port=3306
dbname=douyin2

tablename1=user_favorite_list
tablename2=user_follow_list
#tablename3=user
#tablename4=user_data

username=healer
passwd=healer000.


echo "开始连接库：$dbname 的表：$tablename1"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tablename1}"  -dir="${modeldir}" --style=goZero --cache=true --home=../../../../tpl
echo "开始连接库：$dbname 的表：$tablename2"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tablename2}"  -dir="${modeldir}" --style=goZero --cache=true --home=../../../../tpl
#echo "开始连接库：$dbname 的表：$tablename3"
#goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tablename3}"  -dir="${modeldir}" --style=goZero --cache=false --home=../../../../tpl
#echo "开始连接库：$dbname 的表：$tablename4"
#goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tablename4}"  -dir="${modeldir}" --style=goZero --cache=false --home=../../../../tpl
