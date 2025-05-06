package entity

import "time"

type User struct {
	ID              int       `gorm:"column:id;type:int;primaryKey;autoIncrement:true;unique" json:"id"`
	Name            string    `gorm:"column:name;type:string;size:255;unique" json:"name"`
	Email           string    `gorm:"column:email;type:string;size:255" json:"email"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at;" json:"email_verified_at"`
	Password        string    `gorm:"column:password;type:string;size:255" json:"password"`
	RememberToken   string    `gorm:"column:remember_token;type:string;size:255" json:"remember_token"`

	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`

	// custom
	IsEmpty bool `gorm:"-" json:"-"`
}

func (t User) TableName() string {
	return "users"
}
