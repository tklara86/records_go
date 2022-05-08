package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/tklara86/records_go/internal/models"
)

// homeAdmin
func (app *application) homeAdmin(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/common/navigation.tmpl",
		"./ui/html/pages/common/sidebar.tmpl",
		"./ui/html/pages/admin/about.tmpl",
		"./ui/html/pages/admin/home.tmpl",
	}

	type Link struct {
		LinkTitle string
		LinkPath  string
	}

	type Links struct {
		Link []Link
	}

	data := Links{
		Link: []Link{
			{
				LinkTitle: "Home",
				LinkPath:  "/admin",
			},
			{
				LinkTitle: "Categories",
				LinkPath:  "/admin/categories/",
			},
		},
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) createRecordPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "Actions"
	release := "2016"
	image := "actions.jpg"

	record := &models.Record{
		Title:       title,
		ReleaseDate: release,
		Image:       image,
	}

	id, err := app.records.Insert(record)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// recordGenre := &models.RecordGenre{
	// 	RecordID: int64(id),
	// 	GenreID:  1,
	// }

	// id2, err := app.genres.InsertRecordGenre(recordGenre)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	w.Write([]byte(fmt.Sprintf("Record created with id of %d", id)))
	//w.Write([]byte(fmt.Sprintf("Record Genre created with id of %d", record.RecordGenres.RecordID)))
}

func (app *application) createRecordGet(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/common/navigation.tmpl",
		"./ui/html/pages/common/sidebar.tmpl",
		"./ui/html/pages/record/create.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) viewRecord(w http.ResponseWriter, r *http.Request) {

	// files := []string{
	// 	"./ui/html/pages/base.tmpl",
	// 	"./ui/html/pages/common/navigation.tmpl",
	// 	"./ui/html/pages/common/sidebar.tmpl",
	// 	"./ui/html/pages/admin/about.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)

	// if err != nil {
	// 	app.serverError(w, err)
	// }

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	record, err := app.records.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", record)
}

func (app *application) viewRecords(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }

	records, err := app.records.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, record := range records {
		fmt.Fprintf(w, "%+v\n", record)
	}
}
