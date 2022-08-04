package model

import (
	"time"

	"gorm.io/gorm"
)

type ChatMsg struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	Text       string    `json:"text" gorm:"not null"`
	SenderID   string    `json:"sender_id" gorm:"not null"`
	ChatRoomID string    `json:"chat_room_id" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type slimChatMsg struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SlimChatMsgs []slimChatMsg

func (p *ChatMsg) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}

func (p *ChatMsg) Updates() (tx *gorm.DB) {
	return DB.Model(&p).Updates(p)
}

func (p *ChatMsg) DeleteById(msgID string) (tx *gorm.DB) {
	return DB.Where("id = ?", msgID).Delete(&p)
}

func (p *ChatMsg) FirstById(msgID string) (tx *gorm.DB) {
	return DB.Where("id = ?", msgID).First(&p)
}

func (p *SlimChatMsgs) FetchRoomMsgs(chatRoomID string) (tx *gorm.DB) {
	return DB.Table("chat_msgs").Select("chat_msgs.id, chat_msgs.text, users.name, chat_msgs.created_at, chat_msgs.updated_at").Joins("left join users as users ON users.id = chat_msgs.sender_id").Where("chat_room_id = ?", chatRoomID).Order("created_at asc").Find(&p)
}
