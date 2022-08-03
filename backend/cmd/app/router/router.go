package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Hibagon1go/ChatApp_Go_React/cmd/app/auth"
	"github.com/Hibagon1go/ChatApp_Go_React/cmd/app/controller"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"https://chat-app-go-react-git-chatroom-hibagon1go.vercel.app", "http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}
	e.Use(middleware.CORSWithConfig(corsConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)

	// echo.middleware JWTConfigの設定
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte(auth.JwtKey),
		ParseTokenFunc: func(tokenString string, c echo.Context) (interface{}, error) {
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(auth.JwtKey), nil
			}

			token, err := jwt.Parse(tokenString, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}

	// /api配下のAPIには認証が必要
	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(jwtConfig))
	e.Use(middleware.CORSWithConfig(corsConfig))
	api.GET("/chatroom/all", controller.FetchAllRooms)
	api.GET("/chatroom/joining", controller.FetchJoiningRooms)
	api.POST("/chatroom/build", controller.BuildChatRoom)
	api.POST("/chatroom/join", controller.JoinNewRoom)
	api.PUT("/chatroom/rename", controller.UpdateRoomName)
	api.DELETE("/chatroom/leave", controller.LeaveChatRoom)

	api.GET("/msg", controller.FetchRoomMsgs)
	api.POST("/msg", controller.PublishChatMsg)
	api.PUT("/msg", controller.EditChatMsg)
	api.DELETE("/msg", controller.DeleteChatMsg)

	api.POST("/subroom", controller.SubscribeChatRoom)

	e.Logger.Fatal(e.Start(":8080"))
}
