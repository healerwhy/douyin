package userOpt

import (
	"net/http"

	"douyin/service/api/internal/logic/userOpt"
	"douyin/service/api/internal/svc"
	"douyin/service/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFavoriteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := userOpt.NewGetFavoriteListLogic(r.Context(), svcCtx)
		resp, err := l.GetFavoriteList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
