package main

import (
	"WEB2/learning_http/utils"
	"WEB2/learning_http/views"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

const (
	homeFile    = "views/home.gohtml"
	contactFile = "views/contact.gohtml"
	signupFile  = "views/signup.gohtml"
)

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
	signupTemplate  *template.Template
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeTemplate.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		log.Fatalln("Cannot execute home template", err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactTemplate.ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		log.Fatalln("Cannot execute contact template", err)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactTemplate.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		log.Fatalln("Cannot execute contact template", err)
	}
}

func main() {
	layoutFiles := utils.LayoutFiles()
	homeTemplate = views.ParseTemplates(homeFile, layoutFiles)
	contactTemplate = views.ParseTemplates(contactFile, layoutFiles)
	signupTemplate = views.ParseTemplates(signupFile, layoutFiles)

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signUp", signup)

	err := http.ListenAndServe(":3000", r) // nil - we want use default ServeMux
	if err != nil {
		log.Fatalln("Cannot ListenAndServe, error:", err)
	}
}
