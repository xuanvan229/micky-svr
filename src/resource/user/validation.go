package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"micky-svr/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserModelValidator struct {
	User struct {
		Username string `json:"username" binding:"exists,min=8,max=255"`
		Password string `json:"password" binding:"exists"`
	} `json: "user"`
	UserModel UserModel `json:"-"`
}

func (self *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	fmt.Println(err)
	
	if err != nil {
		return err
	}

	self.UserModel.Username = self.User.Username
	self.UserModel.setPassword(self.User.Password)
	return nil
}

func IsExist(user UserModel, db *gorm.DB) (UserModel, bool) {
	u := UserModel{}
	db.Where("username = ?", user.Username).First(&u)
	if u.Username == user.Username {
		return u, true
	}
	return u, false
}