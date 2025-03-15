package shortlink

import (
	"github.com/lambertstu/shortlink-go/restful/internal/logic/shortlink"
	"github.com/lambertstu/shortlink-go/restful/internal/svc"
	"github.com/lambertstu/shortlink-go/restful/internal/types"
	"github.com/lambertstu/shortlink-go/restful/pkg/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CreateShortLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortLinkCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := shortlink.NewCreateShortLinkLogic(r.Context(), svcCtx)
		resp, err := l.CreateShortLink(&req)
		result.Response(w, resp, err)
	}
}
