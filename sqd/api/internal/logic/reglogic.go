package logic

import (
	"YxEmr/common/xerr"
	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"
	"YxEmr/sqd/rpc/reg/reg"
	"YxEmr/sqd/rpc/reg/reger"
	"context"
	"github.com/pkg/errors"

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

func (l *RegLogic) Reg(req *types.Regreq) (resp *reg.Resp, err error) {
	/// 手动代码开始
	r, err := l.svcCtx.Reger.Do(l.ctx, &reger.Req{
		Ilx:   req.Ilx,
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Csqdh: req.Csqdh,
		Cztbm: req.Cztbm,
	})
	if err != nil {
		return r, errors.Wrapf(xerr.NewErrMsg("登记失败"),
			"登记失败: req: %+v , err : %v ", req, err)
	}

	return r, nil
	//手动代码结束
}
