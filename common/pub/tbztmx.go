package pub

import (
	"YxEmr/common/database"
	"errors"
)

type Tztmx struct {
	CZTBM   string  `gorm:"primary_key;column:CZTBM"`
	IXH     int     `gorm:"primary_key;column:IXH"`
	CZTMC   string  `gorm:"column:CZTMC"`
	CSFXMBM string  `gorm:"column:CSFXMBM"`
	CSFXMMC string  `gorm:"column:CSFXMMC"`
	CDW     string  `gorm:"column:CDW"`
	MDJ     float64 `gorm:"column:MDJ"`
	ICOUNT  float64 `gorm:"column:ICOUNT"`
	MJE     float64 `gorm:"column:MJE"`
	BTJ     bool    `gorm:"column:BTJ"`
	ISFXMZL int     `gorm:"column:ISFXMZL"`
	IDCSF   int     `gorm:"column:IDCSF"`
}

func (u Tztmx) TableName() string {
	return "TBZDZTMX"
}

var cacheZtmxkey = "TZtmx:CZTBM:"

func GetZtmx(cztbm, isfxmzl string) (*[]Tztmx, error) {
	cachekey := cacheZtmxkey + cztbm + ":" + isfxmzl
	var ztmxs []Tztmx
	if database.Cache.IsNotFound(database.Cache.Get(cachekey, &ztmxs)) {
		if err := database.Db.Where("CZTBM = ? AND ISFXMZL <> ?", cztbm, isfxmzl).Find(&ztmxs).Error; err != nil {
			return nil, err
		}
		database.Cache.Set(cachekey, &ztmxs)
	}
	if len(ztmxs) == 0 {
		return nil, errors.New("未查询到相关组套数据")
	}
	return &ztmxs, nil

}
