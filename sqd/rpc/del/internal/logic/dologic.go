package logic

import (
	pub "YxEmr/common"
	"YxEmr/common/database"
	"YxEmr/common/model/struct"
	"YxEmr/sqd/rpc/del/del"
	"YxEmr/sqd/rpc/del/internal/svc"
	"context"
	"errors"
	"strings"

	_ "github.com/jinzhu/copier"
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

func (l *DoLogic) Do(in *del.Req) error {
	db := l.svcCtx.DbEngin
	var (
		tbmx, tbxx, tbxm string
		sqdxx            _struct.Tsqdxx
	)
	switch in.Ibrlx {
	case 0:
		{
			tbxx = database.GetTBName("TBMZJCSQDXXWZX", in.Cbrh)
		}
	case 1:
		{
			tbxx = database.GetTBName("TBZYJCSQDXXWZX", in.Cbrh)
		}
	}

	tbmx = strings.ReplaceAll(tbxx, "XX", "MX")
	tbxm = strings.ReplaceAll(tbxx, "XX", "XM")
	if err := db.Table(tbxx).Where("CBH = ?", in.Csqdh).Find(&sqdxx).Error; err != nil {
		return err
	}
	if (sqdxx == _struct.Tsqdxx{}) {
		return errors.New("未查询到相关申请单数据")
	}
	switch sqdxx.ISFZT {
	case 1, 2:
		{
			return errors.New("申请单已收费，不允许撤销")
		}
	}
	fsql := "Delete From " + tbxx + " Where CBH=" + pub.QuoteStr(in.Csqdh)
	fsql += "Delete From " + tbxm + " Where CBH=" + pub.QuoteStr(in.Csqdh)
	fsql += "Delete From " + tbmx + " Where CBH=" + pub.QuoteStr(in.Csqdh)
	if err := database.Exesql(fsql); err != nil {
		return err
	}
	return nil
}
