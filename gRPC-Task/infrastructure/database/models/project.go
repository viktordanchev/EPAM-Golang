package models

type Project struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
}
