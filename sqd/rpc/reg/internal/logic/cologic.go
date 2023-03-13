package logic

import (
	"YxEmr/common/database"
	"YxEmr/common/pub"
	"context"
	"fmt"

	"YxEmr/sqd/rpc/reg/internal/svc"
	"YxEmr/sqd/rpc/reg/reg"

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

func (l *CoLogic) Co(in *reg.Req) (*reg.Resp, error) {
	var (
		sreg []pub.Treginfo
		regs []interface{}
		fsql string
	)
	db := l.svcCtx.DbEngin
	if err := db.Where("CSQDH = ? ", in.Csqdh).Find(&sreg).Error; err != nil {
		return nil, err
	}
	for _, val := range in.Cztbm {
		fsql += fmt.Sprintf("DELETE FROM TBREGSQDINFO WHERE CSQDH='%s' AND CZTBM='%s'", in.Csqdh, val)

	}
	fsql += database.GetBranchInsertSql(regs, "TBREGSQDINFO")
	//执行登记sql
	if err := database.Exesql(fsql); err != nil {
		return nil, err
	}
	return nil, nil
}
