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

/*func (p *User) EmailAlreadyExists(email string) bool {
	if err := DB.Where("email = ?", email).First(&p).Error; err != nil {
		return false
	} else {
		return true
	}
}

func (p *User) NameAlreadyExists(name string) bool {
	if err := DB.Where("name = ?", name).First(&p).Error; err != nil {
		return false
	} else {
		return true
	}
}*/
