package render

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmplCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var err error
	tmpl := checkTemplateCache(t)
	if tmpl != nil {
		sendTemplateResponse(w, tmpl)
		return
	}

	tmpl, err = createTemplate(t)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	cacheTemplate(t, tmpl)

	sendTemplateResponse(w, tmpl)
}

func checkTemplateCache(t string) *template.Template {
	tmpl, found := tmplCache[t]
	if found {
		return tmpl
	}
	return nil
}

func sendTemplateResponse(w http.ResponseWriter, tmpl *template.Template) {
	tmpl.Execute(w, nil)
}

func createTemplate(t string) (*template.Template, error) {
	path := "./templates/"
	layout := path + "base.layout.html"

	templates := []string{
		layout,
		path + t,
	}

	tmpl, err := template.ParseFiles(templates...)
	return tmpl, err
}

func cacheTemplate(t string, tmpl *template.Template) {
	tmplCache[t] = tmpl
}
