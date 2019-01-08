package user

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"micky-svr/db"
	"log"
	"encoding/json"
)

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


func Register(w http.ResponseWriter, r *http.Request){
		user, err := db.ConnectToCol("blossom_user")
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
		// return c.String(http.StatusOK, "Welcome "+username+"!")
}
