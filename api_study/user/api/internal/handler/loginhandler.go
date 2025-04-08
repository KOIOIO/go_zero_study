package handler

import (
	"go-zeroframework/commn/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zeroframework/api_study/user/api/internal/logic"
	"go-zeroframework/api_study/user/api/internal/svc"
	"go-zeroframework/api_study/user/api/internal/types"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(r, w, resp, err)
	}
}
