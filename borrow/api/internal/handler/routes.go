package handler

import (
	"net/http"

	"book/borrow/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/borrow/do",
				Handler: borrowHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/borrow/return",
				Handler: returnHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
