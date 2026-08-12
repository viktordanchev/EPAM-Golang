package models

type Project struct {
	ProjectId   string `gorm:"column:project_id;primaryKey"`
	Name        string
	Description string
}
