package handler

import (
	"grpc_study/test/api/internal/logic"
	"grpc_study/test/api/internal/svc"
	"grpc_study/test/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewApiLogic(r.Context(), svcCtx)
		resp, err := l.Api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
