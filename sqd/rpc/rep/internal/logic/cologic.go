package logic

import (
	"context"

	"YxEmr/sqd/rpc/rep/internal/svc"
	"YxEmr/sqd/rpc/rep/rep"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoLogic {
	return &CoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CoLogic) Co(in *rep.Req) (*rep.Resp, error) {
	// todo: add your logic here and delete this line

	return &rep.Resp{}, nil
}
