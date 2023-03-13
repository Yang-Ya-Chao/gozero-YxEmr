package logic

import (
	"YxEmr/common/xerr"
	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"
	"YxEmr/sqd/rpc/per/perer"
	"context"
	"github.com/pkg/errors"

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

func (l *PerLogic) Per(req *types.Perreq) (resp *perer.Resp, err error) {
	/// 手动代码开始
	resp, err = l.svcCtx.Perer.Do(l.ctx, &perer.Req{
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Csqdh: req.Csqdh,
		Cztbm: req.Cztbm,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("执行失败"),
			"执行失败: req: %+v , err : %v ", req, err)
	}

	return nil, nil
	// 手动代码结束
}
