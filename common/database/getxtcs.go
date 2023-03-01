package database

import (
	"fmt"
)

type XTCS struct {
	Value string `gorm:"column:cvalue"`
}

var cacheXtcskey = "Txtcs:"

func GetXTCS(name, value interface{}) interface{} {
	var (
		cxtcs XTCS
		ret   interface{}
	)
	key := cacheXtcskey + ":" + name.(string)
	CSQL := fmt.Sprintf(
		"SELECT cvalue From TBYXXTCSI WHERE CSTATUS=1 AND CCSMC='%s' "+
			" union all '"+
			" SELECT cvalue From TBUSERPARAM WHERE CSTATUS=1 AND CNBMC='%s' ",
		name, name)
	if Cache.IsNotFound(Cache.Get(key, &ret)) {
		if err := Db.Raw(CSQL).Scan(&cxtcs).Error; (err != nil) || (cxtcs.Value == "") {
			ret = value
		} else {
			ret = cxtcs.Value
		}
		Cache.Set(key, &ret)
	}
	return ret

}
