package render

import (
	"basic-server/pkg/config"
	"basic-server/pkg/models"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func SetConfig(c *config.AppConfig) {
	app = c
}

// cache compiled prior to run time
func RenderTemplate2(w http.ResponseWriter, t string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
	}

	tmpl, ok := tc[t]
	if !ok {
		log.Fatal("No Template available")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err = tmpl.Execute(buf, &td)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	layoutPattern := "./templates/*.layout.html"
	pagePattern := "./templates/*.page.html"

	tc := map[string]*template.Template{}

	layout, err := filepath.Glob(layoutPattern)
	if err != nil {
		return tc, err
	}

	pages, err := filepath.Glob(pagePattern)
	if err != nil {
		return tc, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl, err := template.ParseFiles(layout[0], page)
		if err != nil {
			return tc, err
		}
		tc[name] = tmpl
	}

	return tc, nil
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
