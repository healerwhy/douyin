package messageTypes

const (
	ActionADD    int64 = 1
	ActionCancel int64 = 2
	ActionErr    int64 = -99
)

// UserFavoriteOptMessage 点赞 / 取消点赞
type UserFavoriteOptMessage struct {
	ActionType int64 `json:"action_type"`
	VideoId    int64 `json:"video_id"`
	UserId     int64 `json:"user_id"`
}

// UserFollowOptMessage 关注 / 取消关注
type UserFollowOptMessage struct {
	ActionType int64 `json:"action_type"`
	FollowId   int64 `json:"follow_id"`
	UserId     int64 `json:"user_id"`
}

// UserCommentOptMessage 评论 / 删除评论
type UserCommentOptMessage struct {
	VideoId     int64  `json:"video_id"`
	CommentId   int64  `json:"comment_id"`
	UserId      int64  `json:"user_id"`
	ActionType  int64  `json:"action_type"`
	CommentText string `json:"comment_text,omitempty"`
	CreateDate  string `json:"create_date,omitempty"`
}
