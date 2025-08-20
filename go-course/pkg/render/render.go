package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/aayushxadhikari/go-course/pkg/config"
	"github.com/aayushxadhikari/go-course/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig



// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}


func AddDefaultData(td *models.TemplateData)(*models.TemplateData){
	return td

}



// RenderTemplate renders a template by name using the template cache
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	}else{
		tc, _  = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from Template Cache")
	}

	buf := new(bytes.Buffer) // hold bytes

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// Creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)
		// Create new template with functions
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Parse layout files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// Cache the template
		myCache[name] = ts
	}

	return myCache, nil
}
