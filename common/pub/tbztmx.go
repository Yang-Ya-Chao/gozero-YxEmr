package pub

import (
	"YxEmr/common/cache"
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

func GetZtmx(cztbm, isfxmzl string) (interface{}, error) {
	cachekey := cacheZtmxkey + cztbm + ":" + isfxmzl
	var ztmxs []Tztmx
	return cache.Take(cachekey, func() (interface{}, error) {
		if err := database.Db.Where("CZTBM = ? AND ISFXMZL <> ?",
			cztbm, isfxmzl).Find(&ztmxs).Error; err != nil {
			return nil, err
		}
		if len(ztmxs) == 0 {
			return nil, errors.New("未查询到相关组套数据")
		}
		return &ztmxs, nil
	})

}
