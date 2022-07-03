package cos

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

type UploaderVideo struct {
	UserId      int64
	MachineId   uint16
	VideoBucket string
	SecretID    string
	SecretKey   string
}

func (l *UploaderVideo) UploadVideo(ctx context.Context, file multipart.File) (string, error) {
	u, _ := url.Parse(l.VideoBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  l.SecretID,
			SecretKey: l.SecretKey,
		},
	})
	genSnowFlake := new(GenSnowFlake)
	id, err := genSnowFlake.GenSnowFlake(l.MachineId)
	if err != nil {
		logx.Errorf("UploadVideo--->GenSnowFlake err : %v", err)
		return "", err
	}
	// 生成useId/id.mp4
	key := strconv.FormatInt(l.UserId, 10) + "/" + strconv.FormatInt(int64(id), 10)
	// 上传视频文件
	_, err = c.Object.Put(ctx, key+".mp4", file, nil)
	if err != nil {
		logx.Errorf("UploadVideo--->Put err : %v", err)
		return "", err
	}

	// 上传成功 返回key
	return key, nil
}
