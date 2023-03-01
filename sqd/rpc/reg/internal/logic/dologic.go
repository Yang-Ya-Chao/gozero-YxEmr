package logic

import (
	pub "YxEmr/common"
	"YxEmr/common/database"
	"YxEmr/common/model/struct"
	"context"
	"errors"
	"fmt"
	"time"

	"YxEmr/sqd/rpc/reg/internal/svc"
	"YxEmr/sqd/rpc/reg/reg"

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
		reg  _struct.Treginfo
		regs []interface{}
		fsql string
	)
	db := l.svcCtx.DbEngin
	for _, val := range in.Cztbm {
		fsql += fmt.Sprintf("DELETE FROM TBREGSQDINFO WHERE CSQDH='%S' AND CZTBM='%S'", in.Csqdh, val)
		if in.Ilx == 1 {
			if err := db.Where("CSQDH = ? AND CZTBM = ?", in.Csqdh, val).Find(&reg).Error; err != nil {
				return err
			}
			if (reg != _struct.Treginfo{}) {
				return errors.New(fmt.Sprintf("申请单[%s]中项目[%s]已被登记", in.Csqdh, val))
			}
			reg = _struct.Treginfo{
				in.Cbrh,
				in.Csqdh,
				int(in.Ibrlx),
				time.Now().Format(pub.TemplateDateTime),
				val,
			}
			regs = append(regs, reg)
		}
	}
	fsql += database.GetBranchInsertSql(regs, "TBREGSQDINFO")
	if err := database.Exesql(fsql); err != nil {
		return err
	}
	return nil
}
