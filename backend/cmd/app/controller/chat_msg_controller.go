package controller

import (
	"net/http"

	"github.com/Hibagon1go/ChatApp_Go_React/cmd/app/auth"
	"github.com/Hibagon1go/ChatApp_Go_React/cmd/app/model"
	"github.com/Hibagon1go/ChatApp_Go_React/cmd/app/redis"
	"github.com/Hibagon1go/ChatApp_Go_React/cmd/app/utils"

	"github.com/labstack/echo/v4"
)

type pubChatMsgReq struct {
	Text       string `json:"text"`
	SenderID   string `json:"sender_id"`
	ChatRoomID string `json:"chat_room_id"`
}

type editChatMsgReq struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type deleteChatMsgReq struct {
	ID string `json:"id"`
}

type subChatRoomReq struct {
	ChatRoomID string `json:"chat_room_id"`
}

func FetchRoomMsgs(c echo.Context) error {
	chatRoomID := c.QueryParam("chat_room_id")
	chatMsgs := model.SlimChatMsgs{}
	chatMsgs.FetchRoomMsgs(chatRoomID)
	return c.JSON(http.StatusOK, chatMsgs)
}

func PublishChatMsg(c echo.Context) error {
	req := new(pubChatMsgReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	msgID := utils.GenerateUUID()
	text := req.Text
	chatRoomID := req.ChatRoomID
	senderID := auth.UserAuth(c)

	chatMsg := model.ChatMsg{
		ID:         msgID,
		Text:       text,
		ChatRoomID: chatRoomID,
		SenderID:   senderID,
	}

	// RedisにPublish
	redis.PublishChatMsg(chatMsg)
	// RDBに永続化
	chatMsg.Create()

	return c.JSON(http.StatusOK, echo.Map{"message": "メッセージの送信に成功しました"})
}

func EditChatMsg(c echo.Context) error {
	req := new(editChatMsgReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	auth.UserAuth(c)

	msgID := req.ID
	newText := req.Text
	targetChatMsg := model.ChatMsg{}
	targetChatMsg.FirstById(msgID)

	chatMsg := model.ChatMsg{
		ID:         msgID,
		Text:       newText,
		ChatRoomID: targetChatMsg.ChatRoomID,
		SenderID:   targetChatMsg.SenderID,
	}

	chatMsg.Updates()

	return c.JSON(http.StatusOK, echo.Map{"message": "メッセージの編集に成功しました"})
}

func DeleteChatMsg(c echo.Context) error {
	req := new(editChatMsgReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	auth.UserAuth(c)

	msgID := req.ID
	targetChatMsg := model.ChatMsg{}
	targetChatMsg.DeleteById(msgID)

	return c.JSON(http.StatusOK, echo.Map{"message": "メッセージの削除に成功しました"})
}

func SubscribeChatRoom(c echo.Context) error {
	req := new(subChatRoomReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	subChatRoomID := req.ChatRoomID
	redis.SubscribeChatRoom(c, subChatRoomID)
	return c.JSON(http.StatusOK, subChatRoomID)
}
