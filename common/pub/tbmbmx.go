package pub

import (
	"YxEmr/common/database"
	"errors"
)

type Tmbmx struct {
	CMBBH    string `gorm:"primary_key;column:CMBBH"`
	CINNERID string `gorm:"primary_key;column:CINNERID"`
	CELEBM   string `gorm:"column:CELEBM"`
	CNAME    string `gorm:"column:CNAME"`
	IXMLX    int    `gorm:"column:IXMLX"`
	CSFXMBM  string `gorm:"column:CSFXMBM"`
	CSFXMMC  string `gorm:"column:CSFXMMC"`
	CJSBT    string `gorm:"column:CJSBT"`
	TGSNR    string `gorm:"column:TGSNR"`
	CKZXKSBM string `gorm:"column:CKZXKSBM"`
	CBGDMBBH string `gorm:"column:CBGDMBBH"`
}

func (u Tmbmx) TableName() string {
	return "TBXMFMBMX"
}

var cacheMbmxkey = "Tmbmx:CMBBH:"

func GetMbmx(key string) (*[]Tmbmx, error) {
	cachekey := cacheMbmxkey + key
	var mbmxs []Tmbmx
	if database.Cache.IsNotFound(database.Cache.Get(cachekey, &mbmxs)) {
		if err := database.Db.Where("CMBBH = ?", key).Find(&mbmxs).Error; err != nil {
			return nil, err
		}
		database.Cache.Set(cachekey, &mbmxs)
	}
	if len(mbmxs) == 0 {
		return nil, errors.New("未查询到相关模板数据")
	}
	return &mbmxs, nil

}
