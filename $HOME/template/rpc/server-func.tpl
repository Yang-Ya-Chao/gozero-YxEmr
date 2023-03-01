
{{if .hasComment}}{{.comment}}{{end}}
func (s *{{.server}}Server) {{.method}} ({{if .notStream}}ctx context.Context,{{if .hasReq}} in {{.request}}{{end}}{{else}}{{if .hasReq}} in {{.request}},{{end}}stream {{.streamBody}}{{end}}) ({{if .notStream}}{{.response}},{{end}}error) {
	l := {{.logicPkg}}.New{{.logicName}}({{if .notStream}}ctx,{{else}}stream.Context(),{{end}}s.svcCtx)
	if resp, err := l.{{.method}}(in); err != nil {
    		l.Logger.Error(err)
    		return &{{.response}}{
    			Code: 0,
    			Msg:  err.Error(),
    			Data: "",
    		}, nil
    	} else {
    		return &{{.response}}{
    			Code: 1,
    			Msg:  "成功",
    			Data: resp,
    		}, nil
    	}

}
