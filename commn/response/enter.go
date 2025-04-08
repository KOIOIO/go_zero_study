package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Response(r *http.Request, w http.ResponseWriter, resp any, err error) {
	if err != nil {
		body := Body{
			Code: 500,
			Data: nil,
			Msg:  err.Error(),
		}
		httpx.WriteJson(w, http.StatusOK, body)
		return
	}
	body := Body{
		Code: 200,
		Data: resp,
		Msg:  "ok",
	}
	httpx.WriteJson(w, http.StatusOK, body)
}
