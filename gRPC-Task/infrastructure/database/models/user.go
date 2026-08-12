package models

type User struct {
	ID           string `gorm:"primaryKey"`
	FirstName    string
	LastName     string
	EmailAddress string
}
