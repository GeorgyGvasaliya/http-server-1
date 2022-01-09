package controllers

import (
	"WEB2/learning_http/views"
	"net/http"
)

const signupFile = "signup.gohtml"

func New(w http.ResponseWriter, r *http.Request) {
	views.NewPage(w, signupFile)
}
