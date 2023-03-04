package database

import (
	"YxEmr/common/cache"
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
	ret, _ = cache.Take(key, func() (interface{}, error) {
		if err := Db.Raw(CSQL).Scan(&cxtcs).Error; (err != nil) || (cxtcs.Value == "") {
			ret = value
		} else {
			ret = cxtcs.Value
		}
		return ret, nil
	})
	return ret

}
