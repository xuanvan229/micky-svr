package user

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"fmt"
	"time"
	"micky-svr/db"
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
	"log"
	"encoding/json"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type User struct {
	Username    string `json : "username"  xml: "username" form: "username" query: "username"`
	Password string `json : "password" xml: "password" form: "password" query: "password"`
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
	
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	user, err := db.ConnectToCol("micky_user")
		if err != nil {
			fmt.Println(err)
	}
	result := new(User)
	err = user.Find(bson.M{"username": username}).One(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if comparePasswords(result.Password, []byte(password)) {
		fmt.Println("true")
		claims := &JwtCustomClaims{
			username,
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			fmt.Println("hello")
		}

		response := map[string]string{"status":t}
		js, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return
	}

	response := map[string]string{"status":"false"}
	js, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(js)
	return
}


func Register(w http.ResponseWriter, r *http.Request){
		user, err := db.ConnectToCol("micky_user")
		if err != nil {
			// return c.String(http.StatusInternalServerError, "false")
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		hashPassword := HashAndSalt([]byte(password))

		newUser := new(User)
		newUser.Username = username
		newUser.Password = hashPassword

		err = user.Insert(newUser)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
			// return c.String(http.StatusInternalServerError, "false")
		}
		
		response := map[string]string{"status":"ok"}
		js, _ := json.Marshal(response)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
		return
}
