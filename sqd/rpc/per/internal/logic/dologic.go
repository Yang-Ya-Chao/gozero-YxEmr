package logic

import (
	"context"

	"YxEmr/sqd/rpc/per/internal/svc"
	"YxEmr/sqd/rpc/per/per"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoLogic {
	return &DoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DoLogic) Do(in *per.Req) (*per.Resp, error) {
	// todo: add your logic here and delete this line

	return &per.Resp{}, nil
}
