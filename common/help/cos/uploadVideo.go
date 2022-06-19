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

type UploaderVideo struct{}

func (*UploaderVideo) UploadVideo(ctx context.Context, file multipart.File, userId int64, machineId uint16, videoBucket, coverBucket, SecretID, SecretKey string) (string, error) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse(videoBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: SecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	genSnowFlake := new(GenSnowFlake)
	id, err := genSnowFlake.GenSnowFlake(machineId)

	// 生成useId/id.mp4
	key := strconv.FormatInt(userId, 10) + "/" + strconv.FormatInt(int64(id), 10)
	// 上传视频文件
	r, err := c.Object.Put(ctx, key+".mp4", file, nil)
	if err != nil {
		logx.Error(r)
		return "", err
	}
	// 上传成功 返回key

	return key, nil
}
