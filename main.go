package main

import (
	"WEB2/learning_http/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/contact", controllers.Contact).Methods("GET")
	r.HandleFunc("/signup", controllers.New).Methods("GET")
	r.HandleFunc("/signup", controllers.Test).Methods("POST")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalln("Cannot ListenAndServe, error:", err)
	}
}
