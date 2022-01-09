package views

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

const (
	layoutsPattern = "views/layouts/*.gohtml"
	viewDir        = "views/"
)

func ParseTemplates(mainFile string, layoutFiles []string) *template.Template {
	tpl, err := template.ParseFiles(append(layoutFiles, viewDir+mainFile)...)
	if err != nil {
		log.Fatalln("Cannot parse template, error:", err)
	}
	return tpl
}

func layoutFiles() []string {
	files, err := filepath.Glob(layoutsPattern)
	if err != nil {
		log.Fatalln(err)
	}
	return files
}

func pageTemplate(pageFile string) *template.Template {
	return ParseTemplates(pageFile, layoutFiles())
}

func NewPage(w http.ResponseWriter, pageFile string) {
	w.Header().Set("Content-Type", "text/html")
	err := pageTemplate(pageFile).ExecuteTemplate(w, pageFile, nil)
	if err != nil {
		log.Fatalln("Cannot execute "+pageFile+" template", err)
	}
}
