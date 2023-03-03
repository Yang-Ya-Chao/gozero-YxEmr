package logic

import (
	"YxEmr/common"
	"YxEmr/common/database"
	"YxEmr/common/pub"
	"YxEmr/sqd/rpc/add/add"
	"YxEmr/sqd/rpc/add/internal/svc"
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
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

func (l *DoLogic) Do(in *add.Req) (string, error) {
	//db := l.svcCtx.DbEngin
	//cache := l.svcCtx.Cache
	var (
		csqdbh, data2, csfxmzl string
		brinfo                 pub.Tbrinfo

		tbmx, tbxx, tbxm string
		sqdmx            pub.Tsqdmx
		sqdmxs           []interface{}
		sqdxm            pub.Tsqdxm
		sqdxms           []interface{}
		sqdxx            pub.Tsqdxx
		sqdxxs           []interface{}
		mbmxs            *[]pub.Tmbmx
		ztmxs            *[]pub.Tztmx
	)
	switch in.Ibrlx {
	case 0:
		{
			tbxx = database.GetTBName("TBMZJCSQDXXWZX", in.Cbrh)
			if mzbr, err := pub.GetMzbr(in.Cbrh); err != nil {
				return "", err
			} else {
				copier.Copy(&brinfo, &mzbr)
				csfxmzl = "1"
				sqdxx.IKSFZT = 3
			}
		}
	case 1:
		{
			if zybr, err := pub.GetZybr(in.Cbrh); err != nil {
				return "", err
			} else {
				copier.Copy(&brinfo, &zybr)
				csfxmzl = "0"
				sqdxx.IKSFZT = 2
				if database.GetXTCS("ZS_SQDXDQZ", 0) == 1 {
					sqdxx.BQZ = true
				}
			}
		}
	}
	tbmx = strings.ReplaceAll(tbxx, "XX", "MX")
	tbxm = strings.ReplaceAll(tbxx, "XX", "XM")
	var err error
	if mbmxs, err = pub.GetMbmx(in.Cmbbh); err != nil {
		return "", err
	}
	if csqdbh, err = database.Getsysnumber("0024", 1, "00"); csqdbh == "" {
		return "", err
	}
	CZTBM := fmt.Sprintf("%+q\n", in.Cztbm)
	for _, u := range *mbmxs {
		switch {
		case strings.Contains(u.CJSBT, "姓名"):
			data2 = brinfo.CXM
		case strings.Contains(u.CJSBT, "性别"):
			data2 = brinfo.CXB
		case strings.Contains(u.CJSBT, "年龄"):
			data2 = brinfo.CNL
		case strings.Contains(u.CJSBT, "住院科室"):
			data2 = brinfo.CKS
		case strings.Contains(u.CJSBT, "门诊科室"):
			data2 = brinfo.CKS
		case strings.Contains(u.CJSBT, "当前操作员"):
			data2 = brinfo.CYS
		case strings.Contains(u.CJSBT, "住院病区"):
			data2 = brinfo.CBQ
		case strings.Contains(u.CJSBT, "地址"):
			data2 = brinfo.CDZ
		case strings.Contains(u.CJSBT, "住址"):
			data2 = brinfo.CDZ
		case strings.Contains(u.CJSBT, "身份证"):
			data2 = brinfo.CSFZH
		case strings.Contains(u.CJSBT, "电话"):
			data2 = brinfo.CDH
		case strings.Contains(u.CJSBT, "门诊号"):
			data2 = in.Cbrh
		case strings.Contains(u.CJSBT, "住院号"):
			data2 = in.Cbrh
		case strings.Contains(u.CJSBT, "床位"):
			data2 = brinfo.CCW
		case strings.Contains(u.CJSBT, "申请时间"):
			data2 = common.Now
		case strings.Contains(u.CJSBT, "挂号时间"):
			data2 = brinfo.DSJ
		case strings.Contains(u.CJSBT, "入院时间"):
			data2 = brinfo.DSJ
		case strings.Contains(u.CJSBT, "申请单号"):
			data2 = csqdbh
		case strings.Contains(u.CJSBT, "申请单编号"):
			data2 = csqdbh
		case strings.Contains(u.CJSBT, "申请医生"):
			data2 = brinfo.CYS
		case strings.Contains(u.CJSBT, "出生时间"):
			data2 = brinfo.DCSNY
		case strings.Contains(u.CJSBT, "出生日期"):
			data2 = brinfo.DCSNY
		default:
			data2 = ""
		}
		if u.CELEBM == "SQD.26" && strings.Contains(CZTBM, "\""+u.CSFXMBM+"\"") {
			data2 = "1"
		}
		sqdmx = pub.Tsqdmx{
			CBH:      csqdbh,
			CINNERID: u.CINNERID,
			CXMBM:    u.CELEBM,
			CDATA1:   u.CSFXMMC,
			Cdata2:   data2,
		}
		sqdmxs = append(sqdmxs, sqdmx)
	}
	//组织项目数据

	var hassfxm = false
	ixh := 1
	CYZNR := ""
	var tmpmbmx pub.Tmbmx
	for _, val := range in.Cztbm {
		hassfxm = false
		for _, u := range *mbmxs {
			hassfxm = u.CSFXMBM == val
			if hassfxm {
				tmpmbmx = u
				sqdxx.CKZXXM += u.CINNERID + "=" + u.CSFXMMC + "|"
				sqdxx.CYZXXM += u.CINNERID + "=" + u.CBGDMBBH + ":0|"
				sqdxx.CBGDBH += u.CINNERID + "=|"
				break
			}
		}
		if !hassfxm {
			return "", errors.New("模板[" + in.Cmbbh + "]与组套项目[" + val + "]未绑定！")
		}
		CYZNR = CYZNR + " " + tmpmbmx.CSFXMMC
		if ztmxs, err = pub.GetZtmx(val, csfxmzl); err != nil {
			return "", err
		}
		for _, u := range *ztmxs {
			sqdxm = pub.Tsqdxm{
				CBH:      csqdbh,
				CINNERID: tmpmbmx.CINNERID,
				CZTBM:    val,
				IXH:      ixh,
				CSFXMBM:  u.CSFXMBM,
				MDJ:      u.MDJ,
				NSL:      u.ICOUNT,
				MCOSTS:   u.MJE,
				MZFJ:     u.MJE,
				ISTATUS:  0,
				CDCSF:    u.IDCSF,
			}
			sqdxms = append(sqdxms, sqdxm)
			ixh += 1
			sqdxx.MCOSTS += u.MJE
			sqdxx.MCOSTSZF += u.MJE
		}
	}
	//组织申请单信息数据
	sqdxx.CBH = csqdbh
	sqdxx.CMBBH = in.Cmbbh
	sqdxx.CBRH = in.Cbrh
	sqdxx.CBRID = brinfo.CBRID
	sqdxx.CBRXM = brinfo.CXM
	sqdxx.CBRXB = brinfo.CXB
	sqdxx.CBRNL = brinfo.CNL
	sqdxx.DJLRQ = common.Now
	sqdxx.CYLH = brinfo.CYLH
	sqdxx.DSJSJ = brinfo.DSJ
	sqdxx.CJLRBM = brinfo.IYS
	sqdxx.CJLRMC = brinfo.CYS
	sqdxx.ISTATUS = 1
	sqdxx.CSQZXDWBM = brinfo.IKS
	sqdxx.CSQZXDWMC = brinfo.CKS
	sqdxx.IHJZT = 0
	sqdxx.IZXZT = 0
	sqdxx.ISFZT = 0
	sqdxx.IBGZT = 0
	sqdxx.IJZZT = 0
	sqdxxs = append(sqdxxs, sqdxx)
	fsql := database.GetBranchInsertSql(sqdxxs, tbxx)
	fsql += database.GetBranchInsertSql(sqdxms, tbxm)
	fsql += database.GetBranchInsertSql(sqdmxs, tbmx)
	if err := database.Exesql(fsql); err != nil {
		return "", err
	}
	return csqdbh, nil
}
