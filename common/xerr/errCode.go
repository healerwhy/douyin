package xerr

// 通用错误码成功返回
const (
	OK  int64 = 0
	ERR int64 = 1
)

// 全局错误码
const (
	SERVER_COMMON_ERROR           int64 = 100001
	REUQEST_PARAM_ERROR           int64 = 100002
	TOKEN_EXPIRE_ERROR            int64 = 100003
	TOKEN_GENERATE_ERROR          int64 = 100004
	DB_ERROR                      int64 = 100005
	DB_UPDATE_AFFECTED_ZERO_ERROR int64 = 100006
	SECRET_ERROR                  int64 = 100007
)
