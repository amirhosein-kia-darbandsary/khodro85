package models

type User struct {
	BaseModel

	Username     string `gorm:"type:varchar(20);unique;not null"`
	FirstName    string `gorm:"type:varchar(15)"`
	LastName     string `gorm:"type:varchar(25)"`
	MobileNumber string `gorm:"type:varchar(11);unique"`
	Email        string `gorm:"type:varchar(64);unique"`
	Password     string `gorm:"type:varchar(64)"`
	Enabled      bool   `gorm:"default:true"`

	UserRoles *[]UserRole
}
