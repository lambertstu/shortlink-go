package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"restful/internal/logic/user"
	"restful/internal/svc"
	"restful/internal/types"
)

func UpsertUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpsertUserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := user.NewUpsertUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UpsertUserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
