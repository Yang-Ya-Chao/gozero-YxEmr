package logic

import (
	"YxEmr/sqd/rpc/per/perer"
	"context"

	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerLogic {
	return &PerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PerLogic) Per(req *types.Perreq) (resp *types.Resp, err error) {
	/// 手动代码开始
	r, err := l.svcCtx.Perer.Do(l.ctx, &perer.Req{
		Ilx:   req.Ilx,
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Csqdh: req.Csqdh,
		Cztbm: req.Cztbm,
	})
	if err != nil {
		return nil, err
	}

	return &types.Resp{
		Data: r.Data,
	}, nil
	// 手动代码结束
}
