//KqMessage
package kqueue

//第三方支付回调更改支付状态通知
type UserFavoriteOptMessage struct {
	OptStatus int64  `json:"payStatus"`
	Opt       string `json:"orderSn"`
}
