package model

import (
	"gorm.io/gorm"
	"time"
)

type ChatRoom struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChatRooms []ChatRoom

func (p *ChatRoom) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}

func (p *ChatRooms) FetchAllRooms() (tx *gorm.DB) {
	return DB.Table("chat_rooms").Select("id, name, created_at, updated_at").Find(&p)
}

func (p *ChatRoom) UpdateName() (tx *gorm.DB) {
	return DB.Updates(&p)
}
