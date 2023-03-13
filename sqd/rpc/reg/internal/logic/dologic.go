package logic

import (
	"YxEmr/common"
	"YxEmr/common/database"
	"YxEmr/common/pub"
	"YxEmr/sqd/rpc/reg/internal/svc"
	"YxEmr/sqd/rpc/reg/reg"
	"context"
	"errors"
	"fmt"
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

func (l *DoLogic) Do(in *reg.Req) (*reg.Resp, error) {
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
		if len(sreg) != 0 {
			for _, s := range sreg {
				if s.CZTBM == val {
					return nil, errors.New(fmt.Sprintf("申请单[%s]中项目[%s]已被登记", in.Csqdh, val))
				}
			}
		}
		reg := pub.Treginfo{
			in.Cbrh,
			in.Csqdh,
			int(in.Ibrlx),
			common.Now,
			val,
		}
		regs = append(regs, reg)
	}
	fsql += database.GetBranchInsertSql(regs, "TBREGSQDINFO")
	//执行登记sql
	if err := database.Exesql(fsql); err != nil {
		return nil, err
	}
	return nil, nil
}
