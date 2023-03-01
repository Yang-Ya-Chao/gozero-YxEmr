package pub

type Tsqdmx struct {
	CBH      string `gorm:"primary_key;column:CBH"`
	CINNERID string `gorm:"primary_key;column:CINNERID"`
	CXMBM    string `gorm:"column:CXMBM"`
	CDATA1   string `gorm:"column:CDATA1"`
	Cdata2   string `gorm:"column:Cdata2"`
}

type Tsqdxm struct {
	CBH      string  `gorm:"primary_key;column:CBH"`
	CINNERID string  `gorm:"primary_key;column:CINNERID"`
	CZTBM    string  `gorm:"column:CZTBM"`
	IXH      int     `gorm:"primary_key;column:IXH"`
	CSFXMBM  string  `gorm:"column:CSFXMBM"`
	MDJ      float64 `gorm:"column:MDJ"`
	NSL      float64 `gorm:"column:NSL"`
	MCOSTS   float64 `gorm:"column:MCOSTS"`
	MZFJ     float64 `gorm:"column:MZFJ"`
	CDJH     string  `gorm:"column:CDJH"`
	ISTATUS  int     `gorm:"column:ISTATUS;default:0"`
	CDCSF    int     `gorm:"column:CDCSF;default:0"`
}

type Tsqdxx struct {
	CBH       string  `gorm:"primary_key;column:CBH"`
	CMBBH     string  `gorm:"column:CMBBH"`
	CBRH      string  `gorm:"column:CBRH"`
	CBRID     string  `gorm:"column:CBRID"`
	CBRXM     string  `gorm:"column:CBRXM"`
	CBRXB     string  `gorm:"column:CBRXB"`
	CBRNL     string  `gorm:"column:CBRNL"`
	DJLRQ     string  `gorm:"column:DJLRQ"`
	DSJSJ     string  `gorm:"column:DSJSJ"`
	CJLRBM    string  `gorm:"column:CJLRBM"`
	CJLRMC    string  `gorm:"column:CJLRMC"`
	ISTATUS   int     `gorm:"column:ISTATUS;default:1"`
	CSQZXDWBM string  `gorm:"column:CSQZXDWBM"`
	CSQZXDWMC string  `gorm:"column:CSQZXDWMC"`
	MCOSTS    float64 `gorm:"column:MCOSTS"`
	MCOSTSZF  float64 `gorm:"column:MCOSTSZF"`
	CKZXXM    string  `gorm:"column:CKZXXM"`
	CYZXXM    string  `gorm:"column:CYZXXM"`
	CBGDBH    string  `gorm:"column:CBGDBH"`
	BQZ       bool    `gorm:"column:BQZ;default:false"`
	DQZ       string  `gorm:"column:DQZ"`
	IHJZT     int     `gorm:"column:IHJZT;default:1"`
	ISFZT     int     `gorm:"column:ISFZT;default:0"`
	IZXZT     int     `gorm:"column:IZXZT;default:0"`
	IBGZT     int     `gorm:"column:IBGZT;default:0"`
	IJZZT     int     `gorm:"column:IJZZT;default:0"`
	IKSFZT    int     `gorm:"column:IKSFZT"`
	CYLH      string  `gorm:"column:CYLH"`
	XMLNR     string  `gorm:"column:XMLNR"`
	BLG       bool    `gorm:"column:BLG;default:false"`
}
