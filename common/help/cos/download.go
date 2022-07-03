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

	u, _ := url.Parse(l.CommentBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  l.SecretID,
			SecretKey: l.SecretKey,
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
