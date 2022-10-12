package usermodel

type User struct {
	ID          int    `json:"id" gorm:"column:id;"`
	Name        string `json:"name" gorm:"column:name;"`
	Addr        string `json:"addr" gorm:"column:addr";`
	Phone       string `json:"phone" gorm:"column:phone";`
	BankAccount string `json:"bankAccount" gorm:"column:bankaccount";`
	Balance     int64  `json:"balance" gorm:"column:balance";`
	AmountUsed  int64  `json:"amountUsed" gorm:"column:amountused";`
}

func (u User) TableName() string {
	return "users"
}
