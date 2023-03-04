package pub

import (
	"YxEmr/common/cache"
	"YxEmr/common/database"
	"errors"
)

type Tbrinfo struct {
	CBRH   string
	CXM    string
	CXB    string
	CNL    string
	CBRID  string
	CYLH   string
	IKS    string
	CKS    string
	CCZYGH string
	IYS    string
	CYS    string
	DSJ    string
	DCSNY  string
	CDZ    string
	CSFZH  string
	CBQ    string
	CCW    string
	CDH    string
}

type Tzybr struct {
	CBRH   string `gorm:"primary_key;column:CZYH"`
	CXM    string `gorm:"column:CXM"`
	CXB    string `gorm:"column:CXB"`
	CNL    string `gorm:"column:CNL"`
	CBRID  string `gorm:"column:CBRID"`
	CYLH   string `gorm:"column:CYLH"`
	IKS    string `gorm:"column:IZYKS"`
	CKS    string `gorm:"column:CZYKS"`
	CCZYGH string `gorm:"column:CCZYGH"`
	IYS    string `gorm:"column:IZYYS"`
	CYS    string `gorm:"column:CZYYS"`
	DSJ    string `gorm:"column:DRYSJ"`
	DCSNY  string `gorm:"column:DCSNY"`
	CDZ    string `gorm:"column:CDZ"`
	CSFZH  string `gorm:"column:CSFZH"`
	CBQ    string `gorm:"column:CZYBQ"`
	CCW    string `gorm:"column:CZYCW"`
}

func (u Tzybr) TableName() string {
	return "VTBZYBR"
}

type Tmzbr struct {
	CBRH   string `gorm:"primary_key;column:CMZH"`
	CXM    string `gorm:"column:CXM"`
	CXB    string `gorm:"column:CXB"`
	CNL    string `gorm:"column:CNL"`
	CBRID  string `gorm:"column:CBRID"`
	CYLH   string `gorm:"column:CYLH"`
	IKS    string `gorm:"column:IKSBM"`
	CKS    string `gorm:"column:CKSMC"`
	CCZYGH string `gorm:"column:CCZYGH"`
	IYS    string `gorm:"column:IYSBM"`
	CYS    string `gorm:"column:CYSMC"`
	DSJ    string `gorm:"column:DGH"`
	DCSNY  string `gorm:"column:DCSNY"`
	CDZ    string `gorm:"column:CDZ"`
	CSFZH  string `gorm:"column:CSFZH"`
	CBQ    string
	CCW    string
}

var cacheMzbrkey = "Tmzbr:CMZH:"

func GetMzbr(key string) (interface{}, error) {
	cachekey := cacheMzbrkey + key
	var mzbr Tmzbr
	return cache.Take(cachekey, func() (interface{}, error) {
		if err := database.Db.Table(database.GetTBName("TBMZGHMX", key)).Where("CMZH = ?",
			key).Find(&mzbr).Error; err != nil {
			return nil, err
		}
		if (mzbr == Tmzbr{}) {
			return nil, errors.New("未找到病人数据")
		}
		return &mzbr, nil
	})

}

var cacheZybrkey = "Tzybr:CZYH:"

func GetZybr(key string) (interface{}, error) {
	cachekey := cacheZybrkey + key
	var zybr Tzybr
	return cache.Take(cachekey, func() (interface{}, error) {
		if err := database.Db.Where("CZYH = ?",
			key).Find(&zybr).Error; err != nil {
			return nil, err
		}
		if (zybr == Tzybr{}) {
			return nil, errors.New("未找到病人数据")
		}
		return &zybr, nil
	})

}
