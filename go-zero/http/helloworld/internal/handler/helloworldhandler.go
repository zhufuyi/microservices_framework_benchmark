package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"helloworld/internal/logic"
	"helloworld/internal/svc"
	"helloworld/internal/types"
)

func HelloworldHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHelloworldLogic(r.Context(), svcCtx)
		resp, err := l.Helloworld(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
