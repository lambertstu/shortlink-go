package shortlink

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"restful/internal/logic/shortlink"
	"restful/internal/svc"
	"restful/internal/types"
)

func PageShortLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortLinkPageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := shortlink.NewPageShortLinkLogic(r.Context(), svcCtx)
		resp, err := l.PageShortLink(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
