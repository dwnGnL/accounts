package models

type AccountServices struct{
	Login string `json:"login"`
	Password string `json:"password"`
	Action Action `json:"action"`
}

type Action struct{
	Type int `json:"type"`
	Services []Services `json:"services"`
}

type Services struct{
	Name string `json:"name"`
	Sum float64 `json:"sum"`
}

type Config struct {
	DbURI string `json:"connectUriDb"`
	LogName string `json:"logName"`
	Port string `json:"port"`
}

type Account struct{
	ID int `gorm:"column:ID;primary_key" json:"ID"`
	Fname string `gorm:"column:Fname" json:"Fname"`
	ApiKey string `gorm:"column:APIKey" json:"-"`
	Balance float64 `gorm:"column:Balance" json:"Balance"`
	LastChanged string `gorm:"column:LastChanged" json:"-"`
	Log int `gorm:"column:Log" json:"-"`
}

func (Account) TableName() string {
	return "accounts"
}

type Operation struct{
	ID int `gorm:"column:Id;primary_key" json:"ID"`
	LastSum float64 `gorm:"column:LastSum"`
	Operations string `gorm:"column:Operation"`
	Date string `gorm:"column:Date"`
	Name string `gorm:"column:Name"`
	Id_account int `gorm:"column:id_account"`
}

func (Operation) TableName() string {
	return "operation"
}