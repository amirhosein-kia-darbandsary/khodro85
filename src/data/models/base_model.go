package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID         int          `gorm:"primaryKey"`
	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone;not_null;autoCreateTime"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`

	CreatedBy  sql.NullInt64 `gorm:"not_null"`
	ModifiedBy sql.NullInt64 `gorm:"null"`
	DeletedBy  sql.NullInt64 `gorm:"null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var UserId = sql.NullInt64{Valid: false}

	if value != nil {
		UserId = sql.NullInt64{Valid: true, Int64: value.(int64)}
	}
	m.CreatedBy = UserId
	m.CreatedAt = time.Now()
	return
}

func (m *BaseModel) BeforeUpdated(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var UserId = sql.NullInt64{Valid: false}

	if value != nil {
		UserId = sql.NullInt64{Int64: value.(int64), Valid: true}
	}
	m.CreatedBy = UserId
	m.ModifiedAt = sql.NullTime{Valid: true, Time: time.Now()}
	return
}
