package models

type Role struct {
	BaseModel
	Name string `gorm:"type:varchar(20);unique;not null"`

	UserRoles *[]UserRole
}
