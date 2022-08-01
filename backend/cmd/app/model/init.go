package model

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load("env/.env.dev")
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DSN")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&ChatRoom{})
	DB.AutoMigrate(&ChatMsg{})
	DB.AutoMigrate(&UserChatRoom{})
}
