package note

import (
	"errors"
	"micky-svr/common"
	"micky-svr/config"
	"micky-svr/resource/user"
	"time"

	"github.com/gin-gonic/gin"
)

type NoteModelValidator struct {
	Note struct {
		UserID  uint   `json:"userid" binding:"exists"`
		Title   string `json:"title" binding:"exists"`
		Content string `json:"content" binding:"exists"`
	} `json:"note"`
	NoteModel NoteModel `json:"-"`
}

func (self *NoteModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)

	if err != nil {
		return err
	}

	username := c.MustGet("username").(string)
	err = self.NoteModel.setUserModel(username)
	if err != nil {
		return err
	}
	_ = self.NoteModel.setCreated()

	self.NoteModel.Title = self.Note.Title
	self.NoteModel.Content = self.Note.Content
	return nil
}

func (note *NoteModel) setUserModel(username string) error {
	db, _ := config.Connect()
	n := user.UserModel{}
	db.Where("username = ?", username).First(&n)
	if n.ID == 0 {
		return errors.New("Cannot find any user")
	}
	note.User = n
	return nil
}

func (note *NoteModel) setCreated() error {
	note.CreateAt = time.Now().Unix()
	return nil
}
