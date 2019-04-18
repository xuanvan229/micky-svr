package note

import (
	"time"
	"micky-svr/resource/user"
	"github.com/jinzhu/gorm"
)

type NoteModel struct {
	ID uint
	CreateAt time.Time
	UserID uint
	User user.UserModel `json:"-" gorm:"foreignkey:UserID" `
	Title string `json:"title"`
	Content string `json:"content"`
}

func NewNoteModelValidator() NoteModelValidator {
	return NoteModelValidator{}
}

func CreateTableUser(db *gorm.DB) {
	check := db.HasTable(&NoteModel{});
	if !check {
		db.CreateTable(&NoteModel{})
	}
}
