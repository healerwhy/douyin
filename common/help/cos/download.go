package cos

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

type DownloadComment struct {
	Key           string
	CommentBucket string
	SecretID      string
	SecretKey     string
}

func (l *DownloadComment) DownloadComment(ctx context.Context) ([]*CommentFile, error) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse(l.CommentBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  l.SecretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: l.SecretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})

	// 多goroutine同时读取文件
	keysCh := make(chan []string, 3)
	var wg sync.WaitGroup
	threadpool := 3
	var allComment []*CommentFile
	for i := 0; i < threadpool; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for keys := range keysCh {
				key := keys[0]
				resp, err := c.Object.Get(ctx, key, nil)
				if err != nil {
					panic(err)
				}
				all, err := ioutil.ReadAll(resp.Body)
				_ = resp.Body.Close()

				if err != nil {
					return
				}
				var ret CommentFile
				_ = json.Unmarshal(all, &ret)

				allComment = append(allComment, &ret)

				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}

	isTruncated := true
	prefix := l.Key // 下载 dir 目录下所有文件

	logx.Infof("-------------%+v", prefix)

	for isTruncated {
		opt := &cos.BucketGetOptions{
			Prefix:       prefix,
			EncodingType: "url", // url编码
		}
		// 列出目录
		v, _, err := c.Bucket.Get(ctx, opt)
		if err != nil {
			fmt.Println(err)
			break
		}
		for _, c := range v.Contents {
			key, _ := cos.DecodeURIComponent(c.Key) //EncodingType: "url"，先对 Key 进行 url decode

			keysCh <- []string{key}
		}
		isTruncated = v.IsTruncated
	}
	close(keysCh)
	wg.Wait()

	return allComment, nil
}
