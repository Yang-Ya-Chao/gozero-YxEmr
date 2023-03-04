package database

import (
	pub "YxEmr/common"
	"YxEmr/common/cache"
	"errors"
	"strconv"
	"strings"
)

type Table struct {
	CBM       string `gorm:"column:CBM"`
	CMC       string `gorm:"column:CMC"`
	ITYPE     int    `gorm:"column:ITYPE"`
	IFQSL     int    `gorm:"column:IFQSL"`
	CDATABASE string `gorm:"column:CDATABASE"`
}

func (u Table) TableName() string {
	return "TBSYSTABLES"
}

var cacheTBNamekey = "Table:"

func GetTBName(tbmc, keyvalue string) string {
	cachekey := cacheTBNamekey + tbmc + ":" + keyvalue
	var result interface{}
	var errs error
	if result, errs = cache.Take(cachekey, func() (interface{}, error) {
		var DateTime pub.TDateTimes
		var table Table
		DateTime.Datetime = keyvalue
		if !DateTime.IsDateTime() && keyvalue != "" {
			DateTime.Datetime = "20" + keyvalue[0:2] + "-" + keyvalue[2:4] + "-01 12:12:12"
		}
		Db.Where("CMC = ?", tbmc).Find(&table)
		if table.CBM == "" {
			return "", errors.New("未找到")
		}
		switch table.ITYPE {
		case 0: //单表
			tbmc = table.CDATABASE + ".dbo." + tbmc
		case 1: //单库年表
			if DateTime.IsDateTime() {
				tbmc = table.CDATABASE + ".dbo." + tbmc + DateTime.GetYear()
			}
		case 4: //年库月表
			if DateTime.IsDateTime() {
				year := DateTime.GetYear()
				YearMonth := DateTime.GetYearMonth()
				tbmc = table.CDATABASE + year + ".dbo." + tbmc + YearMonth
			}
		case 5: //年库年表
			if DateTime.IsDateTime() {
				year := DateTime.GetYear()
				tbmc = table.CDATABASE + year + ".dbo." + tbmc + year
			}
		case 7: //单表分区表
			s := keyvalue[len(keyvalue)-1:]
			if pub.IsInt(s) {
				iqu, _ := strconv.Atoi(s)
				tbmc = table.CDATABASE + ".dbo." + tbmc + "_" + pub.FillZeor(iqu, 2)
			}
		case 8: //年库分区表
			ypre := keyvalue[:2]
			fq := keyvalue[len(keyvalue)-1:]
			if pub.IsInt(ypre) && pub.IsInt(fq) {
				ifq, _ := strconv.Atoi(fq)
				//FQSL,_ := strconv.Atoi(table.IFQSL)
				lastfq := ifq % table.IFQSL
				tbmc = table.CDATABASE + "20" + ypre + ".dbo." + tbmc + "_" + pub.FillZeor(lastfq, 2)
			}
		case 9: //库位表
			if pub.IsInt(keyvalue) {
				if len(keyvalue) == 1 {
					tbmc = table.CDATABASE + ".dbo." + tbmc + "0" + keyvalue
				} else {
					tbmc = table.CDATABASE + ".dbo." + tbmc + keyvalue
				}
			}
		case 10: //病区表
			if strings.ToUpper(keyvalue[len(keyvalue)-2:]) == "BQ" {
				tbmc = table.CDATABASE + ".dbo." + tbmc + keyvalue
			} else {
				tbmc = table.CDATABASE + ".dbo." + tbmc + "BQ" + keyvalue
			}
		case 11: //存储过程
			tbmc = table.CDATABASE + ".dbo." + tbmc
		default:
			return "", errors.New("未找到")
		}
		return tbmc, nil
	}); errs != nil {
		return ""
	}
	return result.(string)
}

//func GetTBName(tbmc, keyvalue string) string {
//	resultTB := tbmc
//	var DateTime pub.TDateTimes
//	var table Table
//	DateTime.Datetime = keyvalue
//	if !DateTime.IsDateTime() && keyvalue != "" {
//		DateTime.Datetime = "20" + keyvalue[0:2] + "-" + keyvalue[2:4] + "-01 12:12:12"
//	}
//	Db.Where("CMC = ?", tbmc).Find(&table)
//	if table.CBM == "" {
//		return ""
//	}
//	switch table.ITYPE {
//	case 0: //单表
//		resultTB = table.CDATABASE + ".dbo." + tbmc
//	case 1: //单库年表
//		if DateTime.IsDateTime() {
//			resultTB = table.CDATABASE + ".dbo." + tbmc + DateTime.GetYear()
//		}
//	case 4: //年库月表
//		if DateTime.IsDateTime() {
//			year := DateTime.GetYear()
//			YearMonth := DateTime.GetYearMonth()
//			resultTB = table.CDATABASE + year + ".dbo." + tbmc + YearMonth
//		}
//	case 5: //年库年表
//		if DateTime.IsDateTime() {
//			year := DateTime.GetYear()
//			resultTB = table.CDATABASE + year + ".dbo." + tbmc + year
//		}
//	case 7: //单表分区表
//		s := keyvalue[len(keyvalue)-1:]
//		if pub.IsInt(s) {
//			iqu, _ := strconv.Atoi(s)
//			resultTB = table.CDATABASE + ".dbo." + tbmc + "_" + pub.FillZeor(iqu, 2)
//		}
//	case 8: //年库分区表
//		ypre := keyvalue[:2]
//		fq := keyvalue[len(keyvalue)-1:]
//		if pub.IsInt(ypre) && pub.IsInt(fq) {
//			ifq, _ := strconv.Atoi(fq)
//			//FQSL,_ := strconv.Atoi(table.IFQSL)
//			lastfq := ifq % table.IFQSL
//			resultTB = table.CDATABASE + "20" + ypre + ".dbo." + tbmc + "_" + pub.FillZeor(lastfq, 2)
//		}
//	case 9: //库位表
//		if pub.IsInt(keyvalue) {
//			if len(keyvalue) == 1 {
//				resultTB = table.CDATABASE + ".dbo." + tbmc + "0" + keyvalue
//			} else {
//				resultTB = table.CDATABASE + ".dbo." + tbmc + keyvalue
//			}
//		}
//	case 10: //病区表
//		if strings.ToUpper(keyvalue[len(keyvalue)-2:]) == "BQ" {
//			resultTB = table.CDATABASE + ".dbo." + tbmc + keyvalue
//		} else {
//			resultTB = table.CDATABASE + ".dbo." + tbmc + "BQ" + keyvalue
//		}
//	case 11: //存储过程
//		resultTB = table.CDATABASE + ".dbo." + tbmc
//	default:
//		resultTB = ""
//	}
//
//	return resultTB
//}
