package shortlink

import (
	"github.com/lambertstu/shortlink-go/restful/internal/logic/shortlink"
	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UpdateShortLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortLinkUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := shortlink.NewUpdateShortLinkLogic(r.Context(), svcCtx)
		resp, err := l.UpdateShortLink(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
