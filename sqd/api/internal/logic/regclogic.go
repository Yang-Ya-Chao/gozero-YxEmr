package logic

import (
	"YxEmr/common/xerr"
	"YxEmr/sqd/rpc/reg/reger"
	"context"
	"github.com/pkg/errors"

	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegcLogic {
	return &RegcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegcLogic) Regc(req *types.Regreq) (resp *reger.Resp, err error) {
	resp, err = l.svcCtx.Reger.Co(l.ctx, &reger.Req{
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Csqdh: req.Csqdh,
		Cztbm: req.Cztbm,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("取消登记失败"),
			"取消登记失败: req: %+v , err : %v ", req, err)
	}

	return nil, nil
	//手动代码结束
}
