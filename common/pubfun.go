package common

import (
	"YxEmr/common/database"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	//各个格式化时间模版
	TemplateDateTime = "2006-01-02 15:04:05" //常规类型
	TemplateDate     = "2006-01-02"          //其他类型
	TemplateTime     = "15:04:05"            //其他类型
)

//通过申请单号获取申请单信息表名
func GetTbSQDXX(ibrlx int64, csqdh, cbrh string) (string, string) {
	tbname := ""
	cbh := ""
	switch ibrlx {
	case 0:
		{
			if strings.Contains(csqdh, "JC") {
				tbname = database.GetTBName("TBMZJCSQDXXWZX", cbrh)
				cbh = csqdh[2:]
			} else {
				tbname = database.GetTBName("TBMZJYSQDXXWZX", cbrh)
				if strings.Contains(csqdh, "JY") {
					cbh = csqdh[2:]
				}
			}

		}
	case 1:
		{
			if strings.Contains(csqdh, "JC") {
				tbname = database.GetTBName("TBZYJCSQDXXWZX", cbrh)
				cbh = csqdh[2:]
			} else {
				tbname = database.GetTBName("TBZYJYSQDXXWZX", cbrh)
				if strings.Contains(csqdh, "JY") {
					cbh = csqdh[2:]
				}
			}
		}
	default:
		return "", ""
	}
	return tbname, cbh
}

// 当前时间 string字符串 "2006-01-02 15:04:05"
var Now = time.Now().Format(TemplateDateTime)

// 时间类型
type TDateTimes struct {
	Datetime string
}

// 时间类型
type TDateTime struct {
	Datetime time.Time
}

/*TDateTimes
  时间字符串：获取时间的 年  年月  年月日
*/
/*判断是否是日期*/
func (r *TDateTimes) IsDateTime() bool {
	_, err := time.Parse("2006-01-02 15:04:05", r.Datetime)
	return err == nil
}
func (r *TDateTimes) GetYear() string {
	//r.Datetime.Year()
	year := ""
	Datetime, err := time.Parse("2006-01-02 15:04:05", r.Datetime)
	if err == nil {
		year = strconv.Itoa(Datetime.Year())
	}
	return year
}
func (r *TDateTimes) GetYearMonth() string {
	//r.Datetime.Year()
	yearmonth := ""
	Datetime, err := time.Parse("2006-01-02 15:04:05", r.Datetime)
	if err == nil {
		year := Datetime.Year()
		month := int(Datetime.Month())
		yearmonth = strconv.Itoa(year) + FillZeor(month, 2)
	}
	return yearmonth
}
func (r *TDateTimes) GetYMD() string {
	//r.Datetime.Year()
	yearmonthday := ""
	Datetime, err := time.Parse("2006-01-02 15:04:05", r.Datetime)
	if err == nil {
		year := Datetime.Year()
		month := int(Datetime.Month())
		day := Datetime.Day()
		yearmonthday = strconv.Itoa(year) + FillZeor(month, 2) + FillZeor(day, 2)
	}
	return yearmonthday
}

/*
功能：补0返回字符串
参数：value：传入值   step：总位数
返回：string
*/
func FillZeor(value, step int) string {
	if step <= 0 {
		step = len(strconv.Itoa(value))
	}
	format := "%0" + strconv.Itoa(step) + "d"
	return fmt.Sprintf(format, value)
}

/*
功能：判断是否是数值
参数：s：字符
返回：bool
*/
func IsNumber(s string) bool {
	// 去除首尾空格
	s = strings.TrimSpace(s)
	for i := 0; i < len(s); i++ {
		// 存在 e 或 E, 判断是否为科学计数法
		if s[i] == 'e' || s[i] == 'E' {
			return IsSciNum(s[:i], s[i+1:])
		}
	}
	// 否则判断是否为整数或小数
	return IsInt(s) || IsDec(s)
}

// 是否为科学计数法
func IsSciNum(num1, num2 string) bool {
	// e 前后字符串长度为0 是错误的
	if len(num1) == 0 || len(num2) == 0 {
		return false
	}
	// e 后面必须是整数，前面可以是整数或小数  4  +
	return (IsInt(num1) || IsDec(num1)) && IsInt(num2)
}

// 判断是否为小数
func IsDec(s string) bool {
	// eg: 11.15, -0.15, +10.15, 3., .15,
	// err: +. 0..
	match1, _ := regexp.MatchString(`^[\+-]?\d*\.\d+$`, s)
	match2, _ := regexp.MatchString(`^[\+-]?\d+\.\d*$`, s)
	return match1 || match2
}

// 判断是否为整数
func IsInt(s string) bool {
	match, _ := regexp.MatchString(`^[\+-]?\d+$`, s)
	return match
}

// 字符前后加上引号
func QuoteStr(value string) string {
	return "'" + value + "'"
}

func BoolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
