package entity

import "time"

type Transaction struct {
	ID              int       `gorm:"column:id;type:int;primaryKey;autoIncrement:true;unique" json:"id"`
	UserId          int       `gorm:"column:customer_id;type:int" json:"customer_id"`
	OrderId         int       `gorm:"column:customer_limit_id;type:int" json:"customer_limit_id"`
	TransactionDate time.Time `gorm:"column:email_verified_at;" json:"email_verified_at"`
	Amount          float64   `gorm:"column:otr_price;type:float" json:"otr_price"`
	PaymentMethod   string    `gorm:"column:password;type:string;size:255" json:"password"`
	Status          float64   `gorm:"column:admin_fee;type:float" json:"admin_fee"`

	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`

	User User `json:"users" gorm:"foreignKey:UserId;references:ID"`

	// custom
	IsEmpty bool `gorm:"-" json:"-"`
}

func (t Transaction) TableName() string {
	return "transactions"
}
