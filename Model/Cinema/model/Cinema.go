package cinemamodel

type Cinema struct {
	ID    int    `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Addr  string `json:"addr" gorm:"column:addr";`
	Phone string `json:"phone" gorm:"column:phone"`
}

func (c Cinema) TableName() string {
	return "cinemas"
}
