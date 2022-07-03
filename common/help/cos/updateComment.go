package cos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"net/url"
)

type UpdateComment struct {
	VideoId   int64 `json:"video_id"`
	CommentId int64 `json:"id"`
	UserId    int64 `json:"user_id"`

	Content    string `json:"content,omitempty" copier:"CommentText"`
	CreateDate string `json:"create_date,omitempty"`
}

type CommentFile struct {
	UserId     int64  `json:"user_id"`
	CommentId  int64  `json:"comment_id"`
	CreateDate string `json:"create_date"`
	Content    string `json:"content" copier:"CommentText"`
}

func (l *UpdateComment) UploadComment(ctx context.Context, CommentBucket, SecretID, SecretKey string) (string, error) {
	u, _ := url.Parse(CommentBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretID,
			SecretKey: SecretKey,
		},
	})

	var commentFile CommentFile
	_ = copier.Copy(&commentFile, l)

	ret, _ := json.Marshal(commentFile)
	buf := bytes.NewBuffer(ret)
	key := "video_id_" + fmt.Sprintf("%d/", l.VideoId) + fmt.Sprintf("user_id_%d_comment_id_%d", l.UserId, l.CommentId) + ".json"

	_, err := c.Object.Put(ctx, key, buf, nil)
	if err != nil {
		logx.Errorf("UploadComment error: %v", err)
		return "", err
	}

	return "", nil
}

func (l *UpdateComment) DeleteComment(ctx context.Context, CommentBucket, SecretID, SecretKey string) (string, error) {
	u, _ := url.Parse(CommentBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretID,
			SecretKey: SecretKey,
		},
	})

	key := "video_id_" + fmt.Sprintf("%d/", l.VideoId) + fmt.Sprintf("user_id_%d_comment_id_%d", l.UserId, l.CommentId) + ".json"

	_, err := c.Object.Delete(ctx, key)
	if err != nil {
		logx.Errorf("UploadComment error: %v", err)
		return "", err
	}

	return "", nil
}
