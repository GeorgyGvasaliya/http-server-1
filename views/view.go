package views

import (
	"log"
	"text/template"
)

func ParseTemplates(mainFile string, layoutFiles []string) *template.Template {
	tpl, err := template.ParseFiles(append(layoutFiles, mainFile)...)
	if err != nil {
		log.Fatalln("Cannot parse template, error:", err)
	}
	return tpl
}
