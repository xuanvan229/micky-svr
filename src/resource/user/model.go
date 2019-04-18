package user

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/dgrijalva/jwt-go"
)

type UserModel struct {
	ID        uint
	CreatedAt time.Time
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func (user *UserModel) setPassword(pw string) error{
	pwd := []byte(pw)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password =string(hash)
	return nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	//fmt.Println(byteHash,plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func createToken(user UserModel) (string, error){
	claims := &JwtCustomClaims{
		user.Username,
		user.Password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return t,err
	}
	return t,nil
}

func NewUserModelValidator() UserModelValidator {
	return UserModelValidator{}
}
func CreateTableUser(db *gorm.DB) {
	check := db.HasTable(&UserModel{});
	if !check {
		db.CreateTable(&UserModel{})
	}
}