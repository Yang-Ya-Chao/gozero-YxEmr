package logic

import (
	"YxEmr/sqd/rpc/cha/chaer"
	"context"

	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChaLogic {
	return &ChaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChaLogic) Cha(req *types.Chareq) (resp *types.Resp, err error) {
	/// 手动代码开始
	r, err := l.svcCtx.Chaer.Do(l.ctx, &chaer.Req{
		Ilx:   req.Ilx,
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Csfr:  req.Csfr,
		Cylkh: req.Cylkh,
		Csqdh: req.Csqdh,
		Cztbm: req.Cztbm,
	})
	if err != nil {
		return nil, err
	}

	return &types.Resp{
		Code: r.Code,
		Msg:  r.Msg,
		Data: r.Data,
	}, nil
	// 手动代码结束
}
