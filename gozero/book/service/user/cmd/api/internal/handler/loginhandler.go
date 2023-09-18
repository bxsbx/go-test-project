package handler

import (
	"net/http"

	"StandardProject/gozero/book/service/user/cmd/api/internal/logic"
	"StandardProject/gozero/book/service/user/cmd/api/internal/svc"
	"StandardProject/gozero/book/service/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
