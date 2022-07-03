package xerr

var message map[int64]string

func init() {
	message = make(map[int64]string)
	message[OK] = "SUCCESS"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[SECRET_ERROR] = "密码错误"
}

func MapErrMsg(errcode int64) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}
