package page

import (
	"micky-svr/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gin-gonic/gin"
	"fmt"
)

type Content struct {
	gorm.Model
	Data string `json:"data"`
	PageId int
}


type Page struct {
	gorm.Model
	Title string `json:"title"`
	Content []Content `gorm:"foreignkey:PageId;" json:"content"`
}


func CreatePage(c *gin.Context) {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Cant not connect db", err)
	}
	check := db.HasTable(&Page{});
	if check {
		pages := []Page{}
		db.Preload("Content").Find(&pages)
		c.JSON(200, pages)
		return
	}
	db.CreateTable(&Page{})
	db.CreateTable(&Content{})
	response := map[string]string{"status": "ok"}
	c.JSON(200, response)
	return
}