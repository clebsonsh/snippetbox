package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/clebsonsh/snippetbox/internal/models"
	"github.com/clebsonsh/snippetbox/ui"
)

type templateData struct {
	Form            any
	Snippet         models.Snippet
	User            models.User
	Flash           string
	CSRFToken       string
	Snippets        []models.Snippet
	CurrentYear     int
	IsAuthenticated bool
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}
