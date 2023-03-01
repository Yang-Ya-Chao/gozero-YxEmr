{{if .hasComment}}{{.comment}}{{end}}
func (l *{{.logicName}}) {{.method}} ({{if .hasReq}}in {{.request}}{{if .stream}},stream {{.streamBody}}{{end}}{{else}}stream {{.streamBody}}{{end}}) (string,error) {
	// todo: add your logic here and delete this line
    db := l.svcCtx.DbEngin
	return "",nil
}
