package handler

import (
	"YxEmr/common/result"
	"net/http"

	"YxEmr/sqd/api/internal/logic"
	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Perreq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewPerLogic(r.Context(), svcCtx)
		resp, err := l.Per(&req)
		result.HttpResult(r, w, resp, err)
	}
}
