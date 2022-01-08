package controllers

import (
	"WEB2/learning_http/views"
	"log"
	"net/http"
	"text/template"
)

const signupFile = "views/signup.gohtml"

func signupTemplate() *template.Template {
	return views.ParseTemplates(signupFile, views.LayoutFiles())
}

func New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := signupTemplate().ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		log.Fatalln("Cannot execute contact template", err)
	}
}
