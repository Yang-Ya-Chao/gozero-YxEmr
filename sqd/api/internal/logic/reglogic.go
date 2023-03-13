package logic

import (
	"YxEmr/common/xerr"
	"YxEmr/sqd/api/internal/svc"
	"YxEmr/sqd/api/internal/types"
	"YxEmr/sqd/rpc/reg/reger"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegLogic {
	return &RegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegLogic) Reg(req *types.Regreq) (resp *reger.Resp, err error) {
	/// 手动代码开始
	in := &reger.Req{
		Ibrlx: req.Ibrlx,
		Cbrh:  req.Cbrh,
		Csqdh: req.Csqdh,
		Cztbm: req.Cztbm,
	}
	//dtm分布式事务
	//dtmServer := l.svcCtx.Config.Dtm
	//reg, _ := l.svcCtx.Config.Reg.BuildTarget()
	//per, _ := l.svcCtx.Config.Per.BuildTarget()
	//gid := shortuuid.New()
	//msg := dtmgrpc.NewMsgGrpc(dtmServer, gid).
	//	Add(reg+"/reg.reger/Do", in).
	//	Add(per+"/per.perer/Do", in)
	//msg.RetryCount = 0
	//msg.WaitResult = true
	//fmt.Println(msg.Gid)
	//err = msg.Submit()
	//直接调用
	resp, err = l.svcCtx.Reger.Do(l.ctx, in)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("登记失败"),
			"登记失败: req: %+v , err : %v ", req, err)
	}
	return nil, nil
	//手动代码结束
}
