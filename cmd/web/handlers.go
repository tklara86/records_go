package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/tklara86/records_go/internal/models"
)

// Home Admin
func (app *application) homeAdmin(w http.ResponseWriter, r *http.Request) {

	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "home.tmpl", data)

}

// recordCreatePost - creates a new record
func (app *application) recordCreatePost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	recordTitle := r.PostForm.Get("recordTitle")
	realeaseDate := r.PostForm.Get("recordReleaseDate")

	image := "actions.jpg"

	record := &models.Record{
		Title:       recordTitle,
		ReleaseDate: realeaseDate,
		Image:       image,
	}

	// New record
	id, err := app.records.Insert(record)
	if err != nil {
		app.serverError(w, err)
		return
	}

	recordGenre := []models.RecordGenre{
		{
			RecordID: int64(id),
			GenreID:  1,
		},
		{
			RecordID: int64(id),
			GenreID:  2,
		},
		{
			RecordID: int64(id),
			GenreID:  3,
		},
	}

	// Genres
	_, err = app.genres.InsertRecordGenre(recordGenre)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write([]byte(fmt.Sprintf("Record created with id of %d", id)))

}

// recordCreateGet - displays a HTML form for creating a new record
func (app *application) recordCreateGet(w http.ResponseWriter, r *http.Request) {

	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "create.tmpl", data)

}

// viewRecord - dipslays single record
func (app *application) viewRecord(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
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

	data := app.newTemplateData(r)
	data.Record = record

	app.render(w, http.StatusOK, "view.tmpl", data)
}

// viewRecords - dipslays all records
func (app *application) viewRecords(w http.ResponseWriter, r *http.Request) {

	records, err := app.records.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Records = records

	app.render(w, http.StatusOK, "records.tmpl", data)

}
