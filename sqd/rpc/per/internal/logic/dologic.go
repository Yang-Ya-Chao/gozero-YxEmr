package logic

import (
	"YxEmr/common"
	"YxEmr/common/pub"
	"context"
	"errors"
	"fmt"
	"strings"

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

func (l *DoLogic) Do(in *per.Req) error {
	// todo: add your logic here and delete this line
	db := l.svcCtx.DbEngin
	var (
		sqdxms                 []pub.Tsqdxm
		tbsqxx, tbsqxm, tbsqmx string
	)
	in.Csqdh, tbsqxx = common.GetTbSQDXX(in.Ibrlx, in.Csqdh, in.Cbrh)

	tbsqxm = strings.ReplaceAll(tbsqxx, "XX", "XM")
	tbsqmx = strings.ReplaceAll(tbsqxx, "XX", "MX")
	fmt.Println(tbsqmx)
	if err := db.Table(tbsqxm).Where("CBH = ?", in.Csqdh).Find(&sqdxms).Error; err != nil {
		return err
	}
	for _, vztbm := range in.Cztbm {
		has := false
		var istatus int
		for _, vsqxm := range sqdxms {
			//查询后续是否有相同组套的不同istatus收费状态
			if has && vztbm == vsqxm.CZTBM {
				if istatus != vsqxm.ISTATUS {
					return errors.New("当前检查项目[" + vztbm + "]收费数据异常")
				}
			}
			//给第一个找到的组套赋值istatus，并且打上has标记
			if !has && vztbm == vsqxm.CZTBM {
				has = true
				istatus = vsqxm.ISTATUS
			}
		}

		if !has {
			return errors.New("申请单未找到检查项目：" + vztbm)
		}

	}
	return nil
}
