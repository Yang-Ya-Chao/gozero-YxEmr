package pub

type Treginfo struct {
	CBRH  string `gorm:"column:CBRH"`
	CSQDH string `gorm:"primary_key;column:CSQDH"`
	IBRLX int    `gorm:"column:IBRLX"`
	DDJRQ string `gorm:"column:DDJRQ"`
	CZTBM string `gorm:"primary_key;column:CZTBM"`
}

func (u Treginfo) TableName() string {
	return "TBREGSQDINFO"
}
