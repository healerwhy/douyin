package feed

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"douyin/service/api/internal/logic/feed"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FeedVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		logx.Errorf("!!!!!!!!!!!!!!--------AuthsInfo-----111111------/n/n/n")

		l := feed.NewFeedVideoListLogic(r.Context(), svcCtx)
		logx.Errorf("!!!!!!!!!!!!!!--------AuthsInfo----2222222-------/n/n/n")

		resp, err := l.FeedVideoList(&req)
		logx.Errorf("!!!!!!!!!!!!!!--------AuthsInfo----333333-------/n/n/n")

		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
