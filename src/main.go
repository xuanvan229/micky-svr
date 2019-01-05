package main

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"fmt"
	"log"
)


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
	r.HandleFunc("/", helloHandler)
	r.HandleFunc("/hello", helloHandler)

	r.Use(loggingMiddleware)
	
	fmt.Println("Sever run at :1323")
	if err := http.ListenAndServe(":1323", r); err != nil {
		panic(err)
	}

}