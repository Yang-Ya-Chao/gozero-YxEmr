package database

import (
	pub "YxEmr/common"
	"fmt"
	"reflect"
	"strings"
)

// GetBranchInsertSql 获取批量添加数据sql语句
func GetBranchInsertSql(objs []interface{}, tableName string) string {
	if len(objs) == 0 {
		return ""
	}
	fieldName := ""
	var valueTypeList []reflect.Kind
	fieldNum := reflect.TypeOf(objs[0]).NumField()
	fieldT := reflect.TypeOf(objs[0])
	for a := 0; a < fieldNum; a++ {
		name := GetColumnName(fieldT.Field(a).Tag.Get("gorm"))
		// 添加字段名
		if a == fieldNum-1 {
			fieldName += fmt.Sprintf("%s", name)
		} else {
			fieldName += fmt.Sprintf("%s,", name)
		}
		// 获取字段类型
		valueTypeList = append(valueTypeList, fieldT.Field(a).Type.Kind())

	}
	var valueList []string
	for _, obj := range objs {
		objV := reflect.ValueOf(obj)
		v := "("
		for index, i := range valueTypeList {
			if index == fieldNum-1 {
				v += GetFormatField(objV, index, i, "")
			} else {
				v += GetFormatField(objV, index, i, ",")
			}
		}
		v += ")"
		valueList = append(valueList, v)
	}
	insertSql := fmt.Sprintf("insert into %s (%s) values %s", tableName, fieldName, strings.Join(valueList, ","))
	return insertSql
}

// 带事务执行原生sql语句
func Exesql(sql string) error {
	tx := Db.Begin()
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// GetFormatField 获取字段类型值转为字符串
func GetFormatField(objV reflect.Value, index int, t reflect.Kind, sep string) string {
	v := ""
	switch t {
	case reflect.String:
		v += fmt.Sprintf("'%s'%s", objV.Field(index).String(), sep)
	case reflect.Bool:
		v += fmt.Sprintf("%d%s", pub.BoolToInt(objV.Field(index).Bool()), sep)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v += fmt.Sprintf("%d%s", objV.Field(index).Int(), sep)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v += fmt.Sprintf("%d%s", objV.Field(index).Uint(), sep)
	case reflect.Float32, reflect.Float64:
		v += fmt.Sprintf("%f%s", objV.Field(index).Float(), sep)
		//case reflect.Interface, reflect.Ptr, reflect.Uintptr:
		//	return "", errors.New(fmt.Sprintf("batch insert unsupport type %s", t.String()))
	}
	return v

}

// GetColumnName 获取字段名
func GetColumnName(jsonName string) string {
	for _, name := range strings.Split(jsonName, ";") {
		if strings.Index(name, "column") == -1 {
			continue
		}
		return strings.Replace(name, "column:", "", 1)
	}
	return ""
}
