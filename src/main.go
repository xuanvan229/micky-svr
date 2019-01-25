package main

import (
	"fmt"
	"io"
	"micky-svr/cms_post"
	"micky-svr/middleware"
	"micky-svr/user"
	"net/http"

	"github.com/gorilla/mux"

	// "encoding/json"
	"log"
	// "io/ioutil"
)

type Login struct {
	Username string `json:"username" form:"username" query:"username"`
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
	api := r.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/hi", helloHandler)
	api.HandleFunc("/login", user.Login).Methods("POST")
	api.HandleFunc("/resgister", user.Register).Methods("POST")
	api.HandleFunc("/check", user.Check).Methods("GET")
	api.HandleFunc("/checkjson", cms_post.CheckJson).Methods("POST")
	p := api.PathPrefix("/admin/").Subrouter()
	p.Use(middleware.LoggingMiddleware)
	p.HandleFunc("/", helloHandler)
	p.HandleFunc("/hello", helloHandler)
	p.HandleFunc("/post", cms_post.GetPost)
	p.HandleFunc("/post/{id}", cms_post.SinglePost)
	fmt.Println("Sever run at :1323")
	if err := http.ListenAndServe(":1323", r); err != nil {
		panic(err)
	}

}
