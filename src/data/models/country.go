package models

type Country struct {
	BaseModel
	ID     int    `gorm:"primarykey"`
	Name   string `gorm:"size:10;type:string;not null;"`
	Cities *[]City
}
