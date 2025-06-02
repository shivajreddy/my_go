package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	CreatedAt    time.Time
	MemberNumber sql.NullString
}

type Model struct {
	ID        uint
	CreatedAt time.Time      `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
