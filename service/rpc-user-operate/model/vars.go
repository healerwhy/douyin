package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const ( // 增加是1 取消是0
	ActionADD    int64 = 1
    ActionCancel int64 = 0
)
