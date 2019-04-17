package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"micky-svr/db"
	"micky-svr/helper"
	"github.com/gin-gonic/gin"
	"micky-svr/common"
	"micky-svr/config"
	"micky-svr/resource/user"
	"errors"
	u "micky-svr/user"
	"net/http"
	"time"
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
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println(r.Method, r.URL)
		cookie, err := r.Cookie("_token")
		startTime := time.Now()
		if err != nil {
			//panic(err)
			helper.SetResponse(&w, "no token", http.StatusForbidden)
			return
		}

		userToken := u.JwtCustomClaims{}
		_, err = jwt.ParseWithClaims(cookie.Value, &userToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			panic(err)
			helper.SetResponse(&w, "false", http.StatusForbidden)
			return
		}

		db, err := sql.Open("postgres", db.DbInfo())

		if err != nil {
			panic(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		db.SetMaxOpenConns(5)
		defer db.Close()
		sqlQuery := `SELECT * FROM mk_user WHERE username=$1 LIMIT 1;`

		row := db.QueryRowContext(ctx, sqlQuery, userToken.Name)
		user := u.User{}
		err = row.Scan(
			&user.Id,
			&user.Username,
			&user.Pass,
		)

		if user.Username == userToken.Name && user.Pass == userToken.Pass {
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
			duration := time.Now().Sub(startTime)
			fmt.Println(duration)
			return
		}
		helper.SetResponse(&w, "false", http.StatusForbidden)
		return

	})
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
			fmt.Println("the code running to here")
			c.Next()
		}
		return 
		// fmt.Println(cookie)
		// c.JSON(200, map[string]string{"status": "ok"}) 
	
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
	// fmt.Println(cookie)
	c.JSON(200, map[string]string{"status": "ok"}) 
}
