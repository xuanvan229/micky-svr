package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"micky-svr/common"
	"micky-svr/config"
	"micky-svr/resource/user"
	"errors"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
)

var ctx = context.Background()
func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods:", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers: Content-Type", "*")
		fmt.Println("go go ")
		next.ServeHTTP(w, r)
		return 
	}) 
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
			if c.Request.Method == "OPTIONS" {
					fmt.Println("OKIE")
					// c.Next()
					c.Abort()
					return
			}

			c.Next()
	}
}

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("_token")
		if err != nil {
			fmt.Println(err)
			c.JSON(404, map[string]string{"status": "not ok 123"})
			c.Abort()
			return 
		}
		
		userToken := user.JwtCustomClaims{}
		_, err = jwt.ParseWithClaims(cookie, &userToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
	
		if err != nil {
			c.JSON(404, map[string]string{"status": "not ok 1234"})
			c.Abort()
			return
		}
	
		userInfo := user.UserModel{}
	
		userInfo.Username = userToken.Username
		userInfo.Password = userToken.Password
	
		db, err := config.Connect()
		if err != nil {
			c.JSON(503, common.ResError("user", err))
			c.Abort()
			return
		}
		
		defer db.Close()
		
		_, isExist := user.IsExist(userInfo, db)
		if !isExist {
			c.JSON(503, common.ResError("user", errors.New("Use does not exist")))
			c.Abort()
			return
		} else {
			c.Set("username", userInfo.Username)
			c.Next()
		}
		return 
	}
}


func CheckLogin(c *gin.Context) {
	cookie, err := c.Cookie("_token")
	if err != nil {
		fmt.Println(err)
		c.JSON(404, map[string]string{"status": "not ok"})
		return 
	}

	userToken := user.JwtCustomClaims{}
	_, err = jwt.ParseWithClaims(cookie, &userToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		c.JSON(404, map[string]string{"status": "not ok"})
		return
	}

	userInfo := user.UserModel{}

	userInfo.Username = userToken.Username
	userInfo.Password = userToken.Password

	db, err := config.Connect()
	if err != nil {
		c.JSON(503, common.ResError("user", err))
		return
	}
	defer db.Close()
	_, isExist := user.IsExist(userInfo, db)
	if !isExist {
		c.JSON(503, common.ResError("user", errors.New("Use does not exist")))
		return
	}
	c.JSON(200, map[string]string{"status": "ok"}) 
}
