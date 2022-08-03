package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var JwtKey string = os.Getenv("JWT_KEY")

func GenerateJWT(userID string) string {
	// ペイロードの作成
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// トークン生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	tokenString, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		return err.Error()
	}
	return tokenString
}

func UserAuth(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	fmt.Print(claims)
	fmt.Print(userID)
	return userID
}
