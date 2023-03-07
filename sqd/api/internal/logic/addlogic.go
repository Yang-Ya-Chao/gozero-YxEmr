package logic

import (
	"YxEmr/common/xerr"
	"YxEmr/sqd/rpc/add/adder"
	"context"
	"github.com/pkg/errors"

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

func (l *AddLogic) Add(req *types.Addreq) (resp string, err error) {
	/// 手动代码开始
	r, err := l.svcCtx.Adder.Do(l.ctx, &adder.Req{
		Isqlx: req.Isqlx,
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Cmbbh: req.Cmbbh,
		Cztbm: req.Cztbm,
	})
	if err != nil {
		logx.Error(err)
		return "", errors.Wrapf(xerr.NewErrMsg("开单失败"),
			"开单失败: req: %+v , err : %v ", req, err)
	}

	return r.Data, nil
	// 手动代码结束
}
