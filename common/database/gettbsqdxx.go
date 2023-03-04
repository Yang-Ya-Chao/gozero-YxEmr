package database

import "strings"

// 通过申请单号获取申请单信息表名 return 表名 申请单号
func GetTbSQDXX(ibrlx int64, csqdh, cbrh string) (string, string) {
	tbname := ""
	cbh := ""
	switch ibrlx {
	case 0:
		{
			if strings.Contains(csqdh, "JC") {
				tbname = GetTBName("TBMZJCSQDXXWZX", cbrh)
				cbh = csqdh[2:]
			} else {
				tbname = GetTBName("TBMZJYSQDXXWZX", cbrh)
				if strings.Contains(csqdh, "JY") {
					cbh = csqdh[2:]
				}
			}

		}
	case 1:
		{
			if strings.Contains(csqdh, "JC") {
				tbname = GetTBName("TBZYJCSQDXXWZX", cbrh)
				cbh = csqdh[2:]
			} else {
				tbname = GetTBName("TBZYJYSQDXXWZX", cbrh)
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
