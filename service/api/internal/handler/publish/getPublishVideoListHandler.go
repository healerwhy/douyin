package publish

import (
	"net/http"

	"douyin/service/api/internal/logic/publish"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPublishVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPubVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := publish.NewGetPublishVideoListLogic(r.Context(), svcCtx)
		resp, err := l.GetPublishVideoList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
