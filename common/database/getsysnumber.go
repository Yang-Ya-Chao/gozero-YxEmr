package database

import (
	"fmt"
)

type Tsysnum struct {
	CBH string `gorm:"column:value"`
}

func Getsysnumber(CBH string, Diff int, TJ string) (string, error) {
	var Sysnum Tsysnum
	// 原生 SQL 查询
	CSQL := fmt.Sprintf("Declare @value VARCHAR(30) SET @value='%s' Exec GetSysNumber2 %d,'%s',@value out Select @value value",
		CBH, Diff, TJ)
	if err := Db.Raw(CSQL).Scan(&Sysnum).Error; err != nil {
		return "", err
	}
	return Sysnum.CBH, nil

}
