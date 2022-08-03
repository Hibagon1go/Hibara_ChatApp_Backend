package controller

import (
	"net/http"

	"github.com/Hibagon1go/ChatApp_Go_React/cmd/app/auth"
	"github.com/Hibagon1go/ChatApp_Go_React/cmd/app/model"
	"github.com/Hibagon1go/ChatApp_Go_React/cmd/app/utils"
	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) error {
	UserID := utils.GenerateUUID()

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	email := user.Email
	password := user.Password
	name := user.Name

	encryptPw, err := auth.PasswordEncrypt(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	newUser := model.User{
		ID:       UserID,
		Email:    email,
		Password: encryptPw,
		Name:     name,
	}

	if !newUser.AlreadyExists(email) {
		newUser.Create()
	}

	token := auth.GenerateJWT(UserID)

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func Login(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	email := user.Email
	password := user.Password

	loginUser := model.User{}
	loginUser.FirstByEmail(email)

	err := auth.CompareHashAndPassword(loginUser.Password, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	token := auth.GenerateJWT(loginUser.ID)

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
