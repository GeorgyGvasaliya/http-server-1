package controllers

import (
	"WEB2/learning_http/views"
	"log"
	"net/http"
	"text/template"
)

const contactFile = "views/contact.gohtml"

func contactTemplate() *template.Template {
	return views.ParseTemplates(contactFile, views.LayoutFiles())
}

func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactTemplate().ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		log.Fatalln("Cannot execute contact template", err)
	}
}
