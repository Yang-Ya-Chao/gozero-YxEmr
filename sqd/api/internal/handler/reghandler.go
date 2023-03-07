package handler

import (
	"YxEmr/common/result"
	"net/http"

	"YxEmr/sqd/api/internal/logic"
	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Regreq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewRegLogic(r.Context(), svcCtx)
		resp, err := l.Reg(&req)
		result.HttpResult(r, w, resp, err)
	}
}
