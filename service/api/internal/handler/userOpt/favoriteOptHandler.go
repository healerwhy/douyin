package userOpt

import (
	"net/http"

	"douyin/service/api/internal/logic/userOpt"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteOptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteOptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := userOpt.NewFavoriteOptLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteOpt(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
