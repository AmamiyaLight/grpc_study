// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4

package handler

import (
	"net/http"

	"grpc_study/video/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/videos/:id",
				Handler: getVideoHandler(serverCtx),
			},
		},
	)
}
