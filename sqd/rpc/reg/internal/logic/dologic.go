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
		fsql += fmt.Sprintf("DELETE FROM TBREGSQDINFO WHERE CSQDH='%s' AND CZTBM='%s'", in.Csqdh, val)
		if in.Ilx == 1 {
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
	}
	fsql += database.GetBranchInsertSql(regs, "TBREGSQDINFO")
	//调用per.rpc执行
	if _, err := l.svcCtx.Perer.Do(l.ctx, &perer.Req{
		Ilx:   in.Ilx,
		Ibrlx: in.Ibrlx,
		Cbrh:  in.Cbrh,
		Csqdh: in.Csqdh,
		Cztbm: in.Cztbm,
	}); err != nil {
		return nil, err
	}
	//dtm分布式事务
	//dtmServer := l.svcCtx.Config.Dtm
	//gid := shortuuid.New()
	//msg := dtmgrpc.NewMsgGrpc(dtmServer, gid).
	//	Add("localhost:9004/per.perer/Do", in)
	//tx := db.Begin()
	//if err := tx.Exec(fsql).Error; err != nil {
	//	tx.Rollback()
	//	return nil, err
	//}
	//if err := msg.Submit(); err != nil {
	//	fmt.Println("per err:", err)
	//	tx.Rollback()
	//}
	//
	//tx.Commit()

	//执行登记sql
	if err := database.Exesql(fsql); err != nil {
		return nil, err
	}

	return nil, nil
}
