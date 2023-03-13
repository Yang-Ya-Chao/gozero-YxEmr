package logic

import (
	"YxEmr/common/xerr"
	"YxEmr/sqd/rpc/cha/chaer"
	"context"
	"github.com/pkg/errors"

	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChacLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChacLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChacLogic {
	return &ChacLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChacLogic) Chac(req *types.Chareq) (resp *chaer.Resp, err error) {
	/// 手动代码开始
	resp, err = l.svcCtx.Chaer.Co(l.ctx, &chaer.Req{
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Csfr:  req.Csfr,
		Cylkh: req.Cylkh,
		Csqdh: req.Csqdh,
		Cztbm: req.Cztbm,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("退费失败"),
			"退费失败: req: %+v , err : %v ", req, err)
	}

	return nil, nil
}
