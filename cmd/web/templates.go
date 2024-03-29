package main

import (
	"html/template"
	"path/filepath"

	"github.com/tklara86/records_go/internal/models"
)

type templateData struct {
	CurrentYear   int
	Links         []Link
	SideBarLinks  []SideBarLink
	FormOptions   []FormOption
	Records       []*models.Record
	Genres        []*models.Genre
	Artists       []*models.Artist
	RecordArtists []*models.Artist
	Labels        []*models.Label
	RecordLabels  *[]models.Label
	Record        *models.Record
	Tracklists    []*models.Tracklist
}

func indexPlusOne(i int) int {
	return i + 1
}

var functions = template.FuncMap{
	"indexPlusOne": indexPlusOne,
}

func newTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/**/*tmpl")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/pages/base.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/pages/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil

}
