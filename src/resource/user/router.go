package user

import (
	"fmt"
	//"fmt"
	"github.com/gin-gonic/gin"
	"micky-svr/config"
	"micky-svr/common"
	"errors"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"time"
)
func UserRegister(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)
}

func Register(c *gin.Context) {

	userModelValidator := NewUserModelValidator()

	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(503, common.ResError("user", err))
		return
	}
	db, err := config.Connect()
	CreateTableUser(db)

	if err != nil {
		c.JSON(503, common.ResError("user", err))
		return
	}

	_, isExist := IsExist(userModelValidator.UserModel, db)
	if isExist {
		c.JSON(503, common.ResError("user", errors.New("User already exist")))
		return
	}

	db.Create(&userModelValidator.UserModel)
	c.JSON(200, userModelValidator.UserModel)
	return
}


func Login(c *gin.Context) {
	userModelValidator := NewUserModelValidator()

	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(503, common.ResError("login", err))
		return
	}
	fmt.Println(userModelValidator.User.Password)
	fmt.Println(userModelValidator.UserModel.Password)

	db, err := config.Connect()
	if err != nil {
		c.JSON(503, common.ResError("user", err))
		return
	}
	user, isExist := IsExist(userModelValidator.UserModel, db)

	if !isExist {
		c.JSON(503, common.ResError("user", errors.New("Use does not exist")))
		return
	}


	if comparePasswords(user.Password, []byte(userModelValidator.User.Password)) {
		token,err := createToken(user)
		if err != nil {
			c.JSON(503, common.ResError("login", err))
		}
		c.SetCookie("_token", token, 3600, "/", "localhost", false, true)
		c.JSON(200, map[string]string{"status": "ok"})
		return
	}
	c.JSON(503, common.ResError("login", errors.New("khong dung password")))
	return
}