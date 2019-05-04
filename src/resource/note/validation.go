package note
import (
	"github.com/gin-gonic/gin"
	"micky-svr/common"
	"micky-svr/resource/user"
	"micky-svr/config"
	"errors"
	"time"
	"fmt"
)

type NoteModelValidator struct {
	Note struct {
		UserID uint `json:"userid" binding:"exists"`
		Title string `json:"title" binding:"exists"`
		Content string `json:"content" binding:"exists"`
	} `json:"note"`
	NoteModel NoteModel `json:"-"`
}


func (self *NoteModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	// fmt.Println(err)
	
	if err != nil {
		return err
	}

	// self.NoteModel.UserID = self.Note.UserID
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

func (note *NoteModel) setUserModel(username string) error{
	db, _ := config.Connect()
	n := user.UserModel{}
	db.Where("username = ?", username).First(&n)
	if(n.ID == 0) {
		return errors.New("Cannot find any user")
	}
	note.User = n
	return nil
}

func (note *NoteModel) setCreated() error{
	fmt.Println("the log of time.now", time.Now())
	note.CreateAt = time.Now().Unix()
	return nil
}