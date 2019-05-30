package note

import (
	"micky-svr/common"
	"micky-svr/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.RouterGroup) {
	router.GET("/note", GetNote)
	router.POST("/note", PostNote)
	router.GET("/note/:id", GetItem)
}

func GetNote(c *gin.Context) {
	listnote, err := common.QueryList(&[]NoteModel{}, c)
	if err != nil {
		c.JSON(503, common.ResError("user", err))
		return
	}

	c.JSON(200, listnote)
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	db, err := config.Connect()
	if err != nil {
		c.JSON(503, common.ResError("user", err))
		return
	}
	defer db.Close()
	note := NoteModel{}
	db.Where("ID = ?", id).First(&note)
	c.JSON(200, note)
	return
}

func PostNote(c *gin.Context) {
	noteModelValidator := NewNoteModelValidator()

	if err := noteModelValidator.Bind(c); err != nil {
		c.JSON(503, common.ResError("user", err))
		return
	}
	db, err := config.Connect()
	CreateTableUser(db)
	defer db.Close()
	if err != nil {
		c.JSON(503, common.ResError("user", err))
		return
	}
	db.Create(&noteModelValidator.NoteModel)
	c.JSON(200, noteModelValidator.NoteModel)
	return
}
