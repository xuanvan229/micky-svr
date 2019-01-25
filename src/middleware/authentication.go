package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"micky-svr/db"
	"micky-svr/helper"
	u "micky-svr/user"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var ctx = context.Background()

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

		fmt.Println("before go to main func =>", time.Now().Sub(startTime))
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
