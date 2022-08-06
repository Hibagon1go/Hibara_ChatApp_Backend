package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *User) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}

func (p *User) FirstByEmail(email string) (tx *gorm.DB) {
	return DB.Where("email = ?", email).First(&p)
}
