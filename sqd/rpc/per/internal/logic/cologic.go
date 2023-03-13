package logic

import (
	"YxEmr/common/database"
	"YxEmr/common/pub"
	"YxEmr/sqd/rpc/per/internal/svc"
	"YxEmr/sqd/rpc/per/per"
	"context"
	"errors"
	"strings"

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

func (l *CoLogic) Co(in *per.Req) (*per.Resp, error) {
	db := l.svcCtx.DbEngin
	var (
		sqdxx          pub.Tsqdxx
		sqdxms         []pub.Tsqdxm
		tbsqxx, tbsqxm string
	)

	//解析申请单号，JC开头的为检查，其他为检验
	tbsqxx, in.Csqdh = database.GetTbSQDXX(in.Ibrlx, in.Csqdh, in.Cbrh)
	if err := db.Table(tbsqxx).Where("CBH = ?", in.Csqdh).Find(&sqdxx).Error; err != nil {
		return nil, err
	}
	if (sqdxx == pub.Tsqdxx{}) {
		return nil, errors.New("未找到申请单信息数据")
	}
	if sqdxx.IZXZT == 0 {
		return nil, nil
	}

	tbsqxm = strings.ReplaceAll(tbsqxx, "XX", "XM")
	if err := db.Table(tbsqxm).Where("CBH = ?", in.Csqdh).Find(&sqdxms).Error; err != nil {
		return nil, err
	}
	if len(sqdxms) == 0 {
		return nil, errors.New("未找到申请单项目数据")
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
					return nil, errors.New("当前检查项目[" + vztbm + "]收费数据异常")
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
			return nil, errors.New("申请单未找到检查项目：" + vztbm)
		}
		if err := db.Where("CELEBM = ? AND CMBBH = ? AND CSFXMBM = ? AND CINNERID = ?",
			"SQD.26", sqdxx.CMBBH, vztbm, sqxm.CINNERID).Find(&mbmx).Error; err != nil {
			return nil, err
		}
		if (mbmx == pub.Tmbmx{}) {
			return nil, errors.New("申请单模板[" + sqdxx.CMBBH + "]未找到对应检查项目[" + vztbm + "]")
		}
		//取消之后的cyzxxm
		cyzxxm0 := sqxm.CINNERID + "=" + mbmx.CBGDMBBH + ":0"
		//执行之后的cyzxxm
		cyzxxm1 := sqxm.CINNERID + "=" + mbmx.CBGDMBBH + ":1"
		if !strings.Contains(sqdxx.CBGDBH, sqxm.CINNERID+"=|") {
			return nil, errors.New("申请单项目[" + vztbm + "]已有报告,禁止取消")
		}
		sqdxx.CYZXXM = strings.Replace(sqdxx.CYZXXM, cyzxxm1, cyzxxm0, -1)
	}
	//更新申请单执行状态--
	//取消执行时，yzxxm中还有:1的数据则为部分取消
	if strings.Contains(sqdxx.CYZXXM, ":1") {
		sqdxx.IZXZT = 2
	} else {
		sqdxx.IZXZT = 0
	}
	if err := db.Table(tbsqxx).Select("*").Updates(&sqdxx).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
