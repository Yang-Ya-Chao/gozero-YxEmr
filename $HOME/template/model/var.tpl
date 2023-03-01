var (
{{.lowerStartCamelObject}}FieldNames = builder.RawFieldNames(&{{.upperStartCamelObject}}{}{{if .postgreSql}}, true{{end}})
{{.lowerStartCamelObject}}Rows = strings.Replace(strings.Join({{.lowerStartCamelObject}}FieldNames, ","),"`","",-1)
{{.lowerStartCamelObject}}RowsExpectAutoSet = strings.Replace({{if .postgreSql}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, {{if .autoIncrement}}"{{.originalPrimaryKey}}", {{end}} {{.ignoreColumns}}), ","){{else}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, {{if .autoIncrement}}"{{.originalPrimaryKey}}", {{end}} {{.ignoreColumns}}), ","){{end}},"`","",-1)
{{.lowerStartCamelObject}}RowsWithPlaceHolder = strings.Replace({{if .postgreSql}}builder.PostgreSqlJoin(stringx.Remove({{.lowerStartCamelObject}}FieldNames, "{{.originalPrimaryKey}}", {{.ignoreColumns}})){{else}}strings.Join(stringx.Remove({{.lowerStartCamelObject}}FieldNames, "{{.originalPrimaryKey}}", {{.ignoreColumns}}), "=?,"),"`","",-1) + "=?"{{end}}

{{if .withCache}}{{.cacheKeys}}{{end}}
)
