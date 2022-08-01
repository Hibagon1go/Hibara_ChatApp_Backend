package redis

import (
	"encoding/json"
	"github.com/Hibagon1go/ChatApp_Go_React/cmd/app/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

// PublishChatMsg ,SubscribeChatRoom : ChatRoomIDをRedis Pub/Subのchannel名として使う
func PublishChatMsg(chatMsg model.ChatMsg) {
	payload, err := json.Marshal(chatMsg)
	if err != nil {
		panic(err)
	}

	if err := redisClient.Publish(ctx, chatMsg.ChatRoomID, payload).Err(); err != nil {
		panic(err)
	}
}

func SubscribeChatRoom(c echo.Context, chatRoomID string) {
	subscriber := redisClient.Subscribe(ctx, chatRoomID)
	chatMsg := model.ChatMsg{}
	ch := subscriber.Channel()
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		for msg := range ch {
			if err := json.Unmarshal([]byte(msg.Payload), &chatMsg); err != nil {
				panic(err)
			}

			websocket.Message.Send(ws, "Received message from "+msg.Channel+" channel")
		}

	}).ServeHTTP(c.Response(), c.Request())
}
