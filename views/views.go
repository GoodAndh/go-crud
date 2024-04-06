package views

import (
	"html/template"
	"net/http"
)

func TemplateExecuted(w http.ResponseWriter, v any, filename ...string) error {
	t, err := template.ParseFiles(filename...)
	if err != nil {
		return err
	}
	return t.Execute(w, v)
}

