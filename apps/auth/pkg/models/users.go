package models

type User struct {
	Code     string `gorm:"primaryKey"`
	Name     string
	LastName string
}

func (user *User) TableName() string {
	return "users"
}
