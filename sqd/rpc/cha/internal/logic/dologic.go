package logic

import (
	"context"

	"YxEmr/sqd/rpc/cha/cha"
	"YxEmr/sqd/rpc/cha/internal/svc"

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

func (l *DoLogic) Do(in *cha.Req) (error) {
	db := l.svcCtx.DbEngin
	db.Debug()

	return nil
}
