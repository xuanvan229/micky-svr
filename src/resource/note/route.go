package note

import (
	"github.com/gin-gonic/gin"
	"micky-svr/common"
	"micky-svr/config"
)

func RegisterRoute(router *gin.RouterGroup) {
	router.GET("/note", GetNote)
	router.POST("/note", PostNote)
}

func GetNote(c *gin.Context) {
	
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