package logic

import (
	"context"

	"YxEmr/sqd/rpc/rep/internal/svc"
	"YxEmr/sqd/rpc/rep/rep"

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

func (l *DoLogic) Do(in *rep.Req) (*rep.Resp, error) {
	// todo: add your logic here and delete this line
	//db := l.svcCtx.DbEngin
	return nil, nil
}
