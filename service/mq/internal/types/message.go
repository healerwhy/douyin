package types

// UserFavoriteOptMessage 点赞 / 取消点赞
type UserFavoriteOptMessage struct {
	OptStatus int64  `json:"payStatus"`
	Opt       string `json:"orderSn"`
}

// UserCommentOptMessage 评论 / 删除评论
type UserCommentOptMessage struct {
	OptStatus int64  `json:"payStatus"`
	Opt       string `json:"orderSn"`
}

// UserFollowOptMessage 关注 / 取消关注
type UserFollowOptMessage struct {
	OptStatus int64  `json:"payStatus"`
	Opt       string `json:"orderSn"`
}
