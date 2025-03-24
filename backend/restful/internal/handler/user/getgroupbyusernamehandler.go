package user

import (
	"github.com/lambertstu/shortlink-go/restful/internal/logic/user"
	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"
	"github.com/lambertstu/shortlink-go/restful/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetGroupByUserNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetGroupRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := user.NewGetGroupByUserNameLogic(r.Context(), svcCtx)
		resp, err := l.GetGroupByUserName(&req)
		result.Response(w, resp, err)
	}
}
