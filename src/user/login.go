package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"micky-svr/db"
	"micky-svr/helper"
	"net/http"
	"time"
	"io/ioutil"
	// "gopkg.in/mgo.v2/bson"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"log"
)

var ctx = context.Background()

type JwtCustomClaims struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
	jwt.StandardClaims
}

type LoginInfo struct {
	Username string `json:"username"  xml:"username" form:"username" query:"username"`
	Password     string `json:"password" xml:"password" form:"password" query:"password"`
}

type User struct {
	Id       int    `json:"id"  xml:"id" form:"id" query:"id"`
	Username string `json:"username"  xml:"username" form:"username" query:"username"`
	Pass     string `json:"pass" xml:"pass" form:"pass" query:"pass"`
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func Login(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("login")
	// if err := r.ParseForm(); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// username := r.FormValue("username")
	// password := r.FormValue("password")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	loginInfo := LoginInfo{}
	err = json.Unmarshal(body, &loginInfo)
	if err != nil {
		panic(err)
	}



	db, err := sql.Open("postgres", db.DbInfo())

	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer db.Close()
	sqlQuery := `SELECT * FROM mk_user WHERE username=$1 LIMIT 1;`

	row := db.QueryRowContext(ctx, sqlQuery, loginInfo.Username)

	user := User{}
	err = row.Scan(
		&user.Id,
		&user.Username,
		&user.Pass,
	)
	if err != nil {
		fmt.Println("get err")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if comparePasswords(user.Pass, []byte(loginInfo.Password)) {
		fmt.Println("true")
		claims := &JwtCustomClaims{
			user.Username,
			user.Pass,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			fmt.Println("hello")
		}

		expire := time.Now()
		cookie := http.Cookie{Name: "_token", Value: t, Path: "/", Expires: expire, MaxAge: 3600}
		http.SetCookie(w, &cookie)
		response := map[string]string{"status": "ok"}
		js, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return
	}
	response := map[string]string{"status": "false"}
	js, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(js)
	return
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	hashPassword := HashAndSalt([]byte(password))

	newUser := new(User)
	newUser.Username = username
	newUser.Pass = hashPassword

	db, err := sql.Open("postgres", db.DbInfo())
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer db.Close()

	sqlQuery := `
	INSERT INTO mk_user (username, pass)
	VALUES ($1, $2);`

	_, err = db.Exec(sqlQuery, newUser.Username, newUser.Pass)

	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]string{"status": "ok"}
	js, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return
}

func Check(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_token")

	if err != nil {
		//panic(err)
		helper.SetResponse(&w, "no token", http.StatusForbidden)
		//response := map[string]string{"status": "no token"}
		//js, _ := json.Marshal(response)
		//w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(http.StatusForbidden)
		//w.Write(js)
		return
	}

	userToken := JwtCustomClaims{}
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

	defer db.Close()
	sqlQuery := `SELECT * FROM mk_user WHERE username=$1 LIMIT 1;`

	row := db.QueryRowContext(ctx, sqlQuery, userToken.Name)

	user := User{}
	err = row.Scan(
		&user.Id,
		&user.Username,
		&user.Pass,
	)

	if user.Username == userToken.Name && user.Pass == userToken.Pass {
		response := map[string]string{"status": "ok"}
		js, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return
	}
	helper.SetResponse(&w, "false", http.StatusForbidden)
	return
}
