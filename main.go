package main

import (
	"WEB2/learning_http/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/contact", controllers.Contact)
	r.HandleFunc("/signup", controllers.New)

	err := http.ListenAndServe(":3000", r) // nil - we want use default ServeMux
	if err != nil {
		log.Fatalln("Cannot ListenAndServe, error:", err)
	}
}
