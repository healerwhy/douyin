package xerr

import (
	"fmt"
)

/**
常用通用固定错误
*/

type CodeError struct {
	errCode int64
	errMsg  string
}

// GetErrCode 返回给前端的错误码
func (e *CodeError) GetErrCode() int64 {
	return e.errCode
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCode(errCode int64) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}
