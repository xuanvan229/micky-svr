package middleware

import (
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"micky-svr/db"
	"micky-svr/helper"
	u "micky-svr/user"
	"net/http"
)

var ctx = context.Background()

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		cookie, err := r.Cookie("_token")

		if err != nil {
			//panic(err)
			helper.FailRequest(&w, "no token", http.StatusForbidden)
			return
		}

		userToken := u.JwtCustomClaims{}
		_, err = jwt.ParseWithClaims(cookie.Value, &userToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			panic(err)
			helper.FailRequest(&w, "false", http.StatusForbidden)
			return
		}

		db, err := sql.Open("postgres", db.DbInfo())

		if err != nil {
			panic(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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
			return
		}
		helper.FailRequest(&w, "false", http.StatusForbidden)
		return

	})
}
