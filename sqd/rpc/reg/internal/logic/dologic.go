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

	//dtmServer := l.svcCtx.Config.Dtm
	//req := &perer.Req{
	//	Ilx:   in.Ilx,
	//	Ibrlx: in.Ibrlx,
	//	Cbrh:  in.Cbrh,
	//	Csqdh: in.Csqdh,
	//	Cztbm: in.Cztbm,
	//}
	//gid := shortuuid.New()
	//msg := dtmgrpc.NewMsgGrpc(dtmServer, gid).
	//	Add("localhost:9004/per.perer/Do", req)

	//dx, _ := db.DB()
	//err := msg.DoAndSubmitDB(fsql, dx, func(tx *sql.Tx) error {
	//	_, err := tx.Exec(fsql)
	//	fmt.Println("4", err)
	//	return err
	//})

	//tx := db.Begin()
	//if err := tx.Exec(fsql).Error; err != nil {
	//	tx.Rollback()
	//	return nil, err
	//}
	//tx.Commit()
	//if err := msg.Submit(); err != nil {
	//	fmt.Println("err:", err)
	//	tx.Rollback()
	//	return nil, err
	//}
	//手动代码结束
	if err := database.Exesql(fsql); err != nil {
		return nil, err
	}

	return nil, nil

	//return nil
}
