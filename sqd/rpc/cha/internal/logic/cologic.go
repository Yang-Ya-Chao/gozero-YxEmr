package logic

import (
	"context"

	"YxEmr/sqd/rpc/cha/cha"
	"YxEmr/sqd/rpc/cha/internal/svc"

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

func (l *CoLogic) Co(in *cha.Req) (*cha.Resp, error) {
	// todo: add your logic here and delete this line

	return &cha.Resp{}, nil
}
