package publish

import (
	"douyin/service/api/internal/logic/publish"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func PublishVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PubVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := publish.NewPublishVideoLogic(r.Context(), svcCtx, r)
		resp, err := l.PublishVideo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
