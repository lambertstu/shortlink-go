package shortlink

import (
	"github.com/lambertstu/shortlink-go/restful/internal/logic/shortlink"
	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"
	"github.com/lambertstu/shortlink-go/restful/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func RestoreUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortLinkRestoreRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := shortlink.NewRestoreUrlLogic(r.Context(), svcCtx)

		resp, err := l.RestoreUrl(&req)
		if err != nil {
			result.FailureWithErr(w, resp, err)
		}
		if resp.OriginUrl == "" {
			result.Failure(w, "源链接不存在或已失效")
		}

		http.Redirect(w, r, resp.OriginUrl, http.StatusFound)
	}
}
