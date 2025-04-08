package handler

import (
	"net/http"

	"go-zeroframework/api_study/user/api_v3/internal/logic"
	"go-zeroframework/api_study/user/api_v3/internal/svc"
	"go-zeroframework/commn/response"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(r, w, resp, err)
	}
}
