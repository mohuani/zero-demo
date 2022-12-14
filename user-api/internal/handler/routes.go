// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	usergroup "zero-demo/user-api/internal/handler/usergroup"
	"zero-demo/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/info",
				Handler: usergroup.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/update",
				Handler: usergroup.UserUpdateHandler(serverCtx),
			},
		},
		rest.WithPrefix("/userapi/v1"),
	)
}
