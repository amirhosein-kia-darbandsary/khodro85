package models

type UserRole struct {
	BaseModel

	UserID uint `gorm:"not null;uniqueIndex:idx_user_role"`
	RoleID uint `gorm:"not null;uniqueIndex:idx_user_role"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
