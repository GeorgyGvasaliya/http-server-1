package controllers

import (
	"WEB2/learning_http/views"
	"net/http"
)

const homeFile = "home.gohtml"

func Home(w http.ResponseWriter, r *http.Request) {
	views.NewPage(w, homeFile)
}
