package views

import (
	"log"
	"path/filepath"
	"text/template"
)

const layoutsPattern = "views/layouts/*.gohtml"

func ParseTemplates(mainFile string, layoutFiles []string) *template.Template {
	tpl, err := template.ParseFiles(append(layoutFiles, mainFile)...)
	if err != nil {
		log.Fatalln("Cannot parse template, error:", err)
	}
	return tpl
}

func LayoutFiles() []string {
	files, err := filepath.Glob(layoutsPattern)
	if err != nil {
		log.Fatalln(err)
	}
	return files
}
