package logic

import (
	"YxEmr/sqd/rpc/reg/reger"
	"context"

	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegLogic {
	return &RegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegLogic) Reg(req *types.Regreq) (resp *types.Resp, err error) {
	/// 手动代码开始
	r, err := l.svcCtx.Reger.Do(l.ctx, &reger.Req{
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
