package {{.pkgName}}

import (
	{{.imports}}
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.logic}} {
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
	/// 手动代码开始
    resp, err = l.svcCtx.xxx.Do(l.ctx,{{.responseType}})
    if err != nil {
    	return nil, errors.Wrapf(xerr.NewErrMsg("失败"),
    		"失败: req: %+v , err : %v ", req, err)
    }

    return resp, nil
}
