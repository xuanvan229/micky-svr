package main

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"fmt"
	"micky-svr/user"
	// "encoding/json"
	"log"
	// "io/ioutil"
)

type Login struct {
	Username  string `json:"username" form:"username" query:"username"`
  	Password string `json:"password" form:"password" query:"password"`
}


func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}



func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/hi", helloHandler)
	r.HandleFunc("/login", user.Login).Methods("POST")
	r.HandleFunc("/resgister", user.Register).Methods("POST")
	p := r.PathPrefix("/admin/").Subrouter()
	p.Use(loggingMiddleware)
	p.HandleFunc("/", helloHandler)
	p.HandleFunc("/hello", helloHandler)

	
	
	fmt.Println("Sever run at :1323")
	if err := http.ListenAndServe(":1323", r); err != nil {
		panic(err)
	}

}