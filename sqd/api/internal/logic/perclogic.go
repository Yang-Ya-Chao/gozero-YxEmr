package logic

import (
	"YxEmr/common/xerr"
	"YxEmr/sqd/rpc/per/perer"
	"context"
	"github.com/pkg/errors"

	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PercLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPercLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PercLogic {
	return &PercLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PercLogic) Perc(req *types.Perreq) (resp *perer.Resp, err error) {
	/// 手动代码开始
	resp, err = l.svcCtx.Perer.Co(l.ctx, &perer.Req{
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Csqdh: req.Csqdh,
		Cztbm: req.Cztbm,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("取消执行失败"),
			"取消执行失败: req: %+v , err : %v ", l.svcCtx.Config.Name, req, err)
	}

	return nil, nil
	// 手动代码结束
}
