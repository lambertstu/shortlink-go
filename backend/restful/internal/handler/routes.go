// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	shortlink "github.com/lambertstu/shortlink-go/restful/internal/handler/shortlink"
	test "github.com/lambertstu/shortlink-go/restful/internal/handler/test"
	user "github.com/lambertstu/shortlink-go/restful/internal/handler/user"
	"github.com/lambertstu/shortlink-go/restful/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/shortlink/create",
				Handler: shortlink.CreateShortLinkHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/shortlink/delete",
				Handler: shortlink.DeleteShortLinkHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/shortlink/page",
				Handler: shortlink.PageShortLinkHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/shortlink/update",
				Handler: shortlink.UpdateShortLinkHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/:shortUri",
				Handler: shortlink.RestoreUrlHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/rabbitmq/send",
				Handler: test.TestRabbitMQHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/group/create",
				Handler: user.CreateGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/group/delete",
				Handler: user.DeleteGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/group/info",
				Handler: user.GetGroupByUserNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/group/update",
				Handler: user.UpdateGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/delete",
				Handler: user.DeleteUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/exist",
				Handler: user.IsExistUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: user.GetUserByUsernameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/login-status",
				Handler: user.IsUserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/logout",
				Handler: user.LogoutHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/upsert",
				Handler: user.UpsertUserInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
