package handler

import (
	"YxEmr/common/result"
	"YxEmr/sqd/api/internal/logic"
	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ChaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Chareq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewChaLogic(r.Context(), svcCtx)
		resp, err := l.Cha(&req)
		result.HttpResult(r, w, resp, err)
	}
}
