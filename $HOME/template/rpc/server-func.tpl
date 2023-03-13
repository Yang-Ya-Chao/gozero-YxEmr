
{{if .hasComment}}{{.comment}}{{end}}
func (s *{{.server}}Server) {{.method}} ({{if .notStream}}ctx context.Context,{{if .hasReq}} in {{.request}}{{end}}{{else}}{{if .hasReq}} in {{.request}},{{end}}stream {{.streamBody}}{{end}}) ({{if .notStream}}{{.response}},{{end}}error) {
	l := {{.logicPkg}}.New{{.logicName}}({{if .notStream}}ctx,{{else}}stream.Context(),{{end}}s.svcCtx)
	r,err := l.{{.method}}(in)
    if err != nil {
    	l.Logger.Error(err)
    }
    return r,err
}
