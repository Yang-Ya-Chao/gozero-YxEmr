package logic

import (
	"YxEmr/common"
	"YxEmr/common/database"
	"YxEmr/common/pub"
	"YxEmr/sqd/rpc/per/perer"
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

func (l *DoLogic) Do(in *reg.Req) error {
	var (
		sreg []pub.Treginfo
		reg  pub.Treginfo
		regs []interface{}
		fsql string
	)
	db := l.svcCtx.DbEngin
	if err := db.Where("CSQDH = ? ", in.Csqdh).Find(&sreg).Error; err != nil {
		return err
	}
	for _, val := range in.Cztbm {
		fsql += fmt.Sprintf("DELETE FROM TBREGSQDINFO WHERE CSQDH='%S' AND CZTBM='%S'", in.Csqdh, val)
		if in.Ilx == 1 {
			if len(sreg) != 0 {
				for _, s := range sreg {
					if s.CZTBM == val {
						return errors.New(fmt.Sprintf("申请单[%s]中项目[%s]已被登记", in.Csqdh, val))
					}
				}
			}
			reg = pub.Treginfo{
				in.Cbrh,
				in.Csqdh,
				int(in.Ibrlx),
				common.Now,
				val,
			}
			regs = append(regs, reg)
		}
	}
	fsql += database.GetBranchInsertSql(regs, "TBREGSQDINFO")
	//调用per.rpc执行
	r, err := l.svcCtx.Perer.Do(l.ctx, &perer.Req{
		Ilx:   in.Ilx,
		Ibrlx: in.Ibrlx,
		Cbrh:  in.Cbrh,
		Csqdh: in.Csqdh,
		Cztbm: in.Cztbm,
	})
	if err != nil {
		return err
	}
	if r.Code == 0 {
		return errors.New(r.Msg)
	}
	// 手动代码结束
	if err := database.Exesql(fsql); err != nil {
		return err
	}
	return nil
}
