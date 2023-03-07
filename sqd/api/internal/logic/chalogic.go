package logic

import (
	"YxEmr/common/xerr"
	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"
	"YxEmr/sqd/rpc/cha/cha"
	"YxEmr/sqd/rpc/cha/chaer"
	"context"
	"github.com/pkg/errors"

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

func (l *ChaLogic) Cha(req *types.Chareq) (resp *cha.Resp, err error) {
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
		return r, errors.Wrapf(xerr.NewErrMsg("计费失败"),
			"计费失败: req: %+v , err : %v ", req, err)
	}

	return r, nil
}
