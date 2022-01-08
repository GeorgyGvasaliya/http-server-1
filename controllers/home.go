package controllers

import (
	"WEB2/learning_http/views"
	"log"
	"net/http"
	"text/template"
)

const homeFile = "views/home.gohtml"

func homeTemplate() *template.Template {
	return views.ParseTemplates(homeFile, views.LayoutFiles())
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeTemplate().ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		log.Fatalln("Cannot execute home template", err)
	}
}
