package logic

import (
	"YxEmr/common/xerr"
	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"
	"YxEmr/sqd/rpc/del/del"
	"YxEmr/sqd/rpc/del/deler"
	"context"
	"github.com/pkg/errors"

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

func (l *DelLogic) Del(req *types.Delreq) (resp *del.Resp, err error) {
	/// 手动代码开始
	r, err := l.svcCtx.Deler.Do(l.ctx, &deler.Req{
		Isqlx: req.Isqlx,
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Csqdh: req.Csqdh,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("撤单失败"),
			"撤单失败: req: %+v , err : %v ", req, err)
	}

	return r, nil
}
