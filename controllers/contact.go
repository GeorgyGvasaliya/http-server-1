package controllers

import (
	"WEB2/learning_http/views"
	"net/http"
)

const contactFile = "contact.gohtml"

func Contact(w http.ResponseWriter, r *http.Request) {
	views.NewPage(w, contactFile)
}
