package logic

import (
	"YxEmr/sqd/rpc/add/adder"
	"context"

	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.Addreq) (resp *types.Resp, err error) {
	/// 手动代码开始
	r, err := l.svcCtx.Adder.Do(l.ctx, &adder.Req{
		Isqlx: req.Isqlx,
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Cmbbh: req.Cmbbh,
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
