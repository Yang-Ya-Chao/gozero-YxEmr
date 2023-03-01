package logic

import (
	"YxEmr/sqd/rpc/del/deler"
	"context"

	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.Delreq) (resp *types.Resp, err error) {
	/// 手动代码开始
	r, err := l.svcCtx.Deler.Do(l.ctx, &deler.Req{
		Isqlx: req.Isqlx,
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Csqdh: req.Csqdh,
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
