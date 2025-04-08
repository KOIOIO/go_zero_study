package handler

import (
	"go-zeroframework/commn/response"
	"net/http"

	"go-zeroframework/api_study/user/api/internal/logic"
	"go-zeroframework/api_study/user/api/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(r, w, resp, err)
	}
}
