package logic

import (
	"context"

	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"

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

func (l *PerLogic) Per(req *types.Perreq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	return
}
