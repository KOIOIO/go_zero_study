package handler

import (
	"go-zeroframework/commn/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zeroframework/rpc_study/user_api_rpc/api/internal/logic"
	"go-zeroframework/rpc_study/user_api_rpc/api/internal/svc"
	"go-zeroframework/rpc_study/user_api_rpc/api/internal/types"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		response.Response(r, w, resp, err)
	}
}
