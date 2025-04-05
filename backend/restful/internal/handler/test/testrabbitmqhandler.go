package test

import (
	"github.com/lambertstu/shortlink-go/restful/internal/logic/test"
	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"
	"github.com/lambertstu/shortlink-go/restful/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func TestRabbitMQHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RabbitMQTestRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := test.NewTestRabbitMQLogic(r.Context(), svcCtx)
		resp, err := l.TestRabbitMQ(&req)
		result.Response(w, resp, err)
	}
}
