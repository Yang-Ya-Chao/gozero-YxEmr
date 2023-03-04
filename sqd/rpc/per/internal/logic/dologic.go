package logic

import (
	"YxEmr/common/database"
	"YxEmr/common/pub"
	"context"
	"errors"
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
	db := l.svcCtx.DbEngin
	var (
		sqdxx          pub.Tsqdxx
		sqdxms         []pub.Tsqdxm
		tbsqxx, tbsqxm string
	)
	//解析申请单号，JC开头的为检查，其他为检验
	tbsqxx, in.Csqdh = database.GetTbSQDXX(in.Ibrlx, in.Csqdh, in.Cbrh)
	if err := db.Table(tbsqxx).Where("CBH = ?", in.Csqdh).Find(&sqdxx).Error; err != nil {
		return err
	}
	if (sqdxx == pub.Tsqdxx{}) {
		return errors.New("未找到申请单信息数据")
	}

	switch {
	//IZXZT--0:已经全部取消，1:已经全部执行 ，此时无需在执行直接返回成功即可
	case sqdxx.IZXZT == int(in.Ilx):
		{
			return nil
		}
	case sqdxx.ISFZT == 0 && in.Ilx == 1:
		{
			return errors.New("申请单未收费,禁止执行")
		}
	case sqdxx.ISFZT == 3 && in.Ilx == 1:
		{
			return errors.New("申请单已退费,禁止执行")
		}
	case sqdxx.IZXZT == 3 && in.Ilx == 1:
		{
			return errors.New("申请单不执行,禁止执行")
		}
	case sqdxx.IZXZT == 4 && in.Ilx == 1:
		{
			return errors.New("申请单已撤销,禁止执行")
		}
	case in.Ibrlx == 1 && !sqdxx.BQZ && in.Ilx == 1:
		{
			return errors.New("申请单未签字,禁止执行")
		}
	}
	tbsqxm = strings.ReplaceAll(tbsqxx, "XX", "XM")
	if err := db.Table(tbsqxm).Where("CBH = ?", in.Csqdh).Find(&sqdxms).Error; err != nil {
		return err
	}
	if len(sqdxms) == 0 {
		return errors.New("未找到申请单项目数据")
	}
	for _, vztbm := range in.Cztbm {
		has := false
		var (
			istatus int
			mbmx    pub.Tmbmx
			sqxm    pub.Tsqdxm
		)
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
				sqxm = vsqxm
			}
		}
		if !has {
			return errors.New("申请单未找到检查项目：" + vztbm)
		}
		if err := db.Where("CELEBM = ? AND CMBBH = ? AND CSFXMBM = ? AND CINNERID = ?",
			"SQD.26", sqdxx.CMBBH, vztbm, sqxm.CINNERID).Find(&mbmx).Error; err != nil {
			return err
		}
		if (mbmx == pub.Tmbmx{}) {
			return errors.New("申请单模板[" + sqdxx.CMBBH + "]未找到对应检查项目[" + vztbm + "]")
		}
		//取消之后的cyzxxm
		cyzxxm0 := sqxm.CINNERID + "=" + mbmx.CBGDMBBH + ":0"
		//执行之后的cyzxxm
		cyzxxm1 := sqxm.CINNERID + "=" + mbmx.CBGDMBBH + ":1"
		switch in.Ilx {
		case 0:
			{ //取消
				if !strings.Contains(sqdxx.CBGDBH, sqxm.CINNERID+"=|") {
					return errors.New("申请单项目[" + vztbm + "]已有报告,禁止取消")
				}
				sqdxx.CYZXXM = strings.Replace(sqdxx.CYZXXM, cyzxxm1, cyzxxm0, -1)
			}
		case 1:
			{ //执行
				if istatus != 2 {
					return errors.New("申请单项目[" + vztbm + "]未收费,禁止执行")
				}
				sqdxx.CYZXXM = strings.Replace(sqdxx.CYZXXM, cyzxxm0, cyzxxm1, -1)
			}
		default:
			return errors.New("操作类型错误")
		}
	}
	//更新申请单执行状态--
	switch in.Ilx {
	case 0:
		{ //取消执行时，yzxxm中还有:1的数据则为部分取消
			if strings.Contains(sqdxx.CYZXXM, ":1") {
				sqdxx.IZXZT = 2
			} else {
				sqdxx.IZXZT = 0
			}
		}
	case 1:
		{ //执行时，yzxxm中还有:0的数据则为部分执行
			if strings.Contains(sqdxx.CYZXXM, ":0") {
				sqdxx.IZXZT = 2
			} else {
				sqdxx.IZXZT = 1
			}
		}
	}
	if err := db.Table(tbsqxx).Select("*").Updates(&sqdxx).Error; err != nil {
		return err
	}
	return nil
}
