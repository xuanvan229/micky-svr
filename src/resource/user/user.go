package user

import (
	"micky-svr/config"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"github.com/gin-gonic/gin"
	"fmt"
)

type User struct {
	ID        uint
	CreatedAt time.Time
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Cant not connect db", err)
	}

	check := db.HasTable(&User{});
	if !check {
		db.CreateTable(&User{})
	}

	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(503, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(200, user)
	return
}