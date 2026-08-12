package models

type User struct {
	UserId       string `gorm:"column:user_id;primaryKey"`
	FirstName    string
	LastName     string
	EmailAddress string
}
