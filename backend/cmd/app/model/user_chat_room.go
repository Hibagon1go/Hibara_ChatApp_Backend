package model

import (
	"gorm.io/gorm"
	"time"
)

type UserChatRoom struct {
	UserID     string    `json:"user_id" gorm:"not null"`
	ChatRoomID string    `json:"chat_room_id" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type JoiningRoom struct {
	ChatRoomID string `json:"chat_room_id"`
	Name       string `json:"name"`
}

type JoiningRooms []JoiningRoom

func (p *UserChatRoom) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}

func (p *UserChatRoom) HasAlreadyJoined(userID string, chatRoomID string) bool {
	if err := DB.Where("user_id = ? and chat_room_id = ?", userID, chatRoomID).First(&p).Error; err != nil {
		return false
	} else {
		return true
	}
}

// FetchJoiningRooms :userIDのユーザが参加しているチャットルーム一覧をとってくる
func (p *JoiningRooms) FetchJoiningRooms(userID string) (tx *gorm.DB) {
	return DB.Table("user_chat_rooms").Select("user_chat_rooms.chat_room_id, chat_rooms.name").Joins("left join chat_rooms as chat_rooms ON chat_rooms.id = user_chat_rooms.chat_room_id").Where("user_id = ?", userID).Find(&p)
}

func (p *UserChatRoom) LeaveRoom(userID string, chatRoomID string) (tx *gorm.DB) {
	return DB.Where("user_id = ? and chat_room_id = ?", userID, chatRoomID).Delete(&p)
}
