package main

import (
	"WEB2/learning http/utils"
	"WEB2/learning http/views"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

const (
	homeFile    = "views/home.gohtml"
	contactFile = "views/contact.gohtml"
)

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
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

func main() {
	layoutFiles := utils.LayoutFiles()
	homeTemplate = views.ParseTemplates(homeFile, layoutFiles)
	contactTemplate = views.ParseTemplates(contactFile, layoutFiles)

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	err := http.ListenAndServe(":3000", r) // nil - we want use default ServeMux
	if err != nil {
		log.Fatalln("Cannot ListenAndServe, error:", err)
	}
}
