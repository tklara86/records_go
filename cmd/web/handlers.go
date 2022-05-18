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

	recordGenre := []models.RecordGenre{}
	for _, g := range r.PostForm["genre-name"] {

		genreID, err := strconv.ParseInt(g, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		rg := []models.RecordGenre{
			{
				RecordID: int64(id),
				GenreID:  genreID,
			},
		}
		recordGenre = append(recordGenre, rg...)
	}

	// RecordGenres
	_, err = app.genres.InsertRecordGenre(recordGenre)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// RecordArtists
	recordArtist := []models.RecordArtist{}
	for _, a := range r.PostForm["artist-name"] {
		artistID, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		ra := []models.RecordArtist{
			{
				ArtistID: artistID,
				RecordID: int64(id),
			},
		}
		recordArtist = append(recordArtist, ra...)
	}

	_, err = app.artists.InsertRecordArtist(recordArtist)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write([]byte(fmt.Sprintf("Record created with id of %d", id)))

}

// recordCreateGet - displays a HTML form for creating a new record
func (app *application) recordCreateGet(w http.ResponseWriter, r *http.Request) {
	// pass genres struct
	genres, err := app.genres.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// pass artists struct
	artists, err := app.artists.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)

	data.Genres = genres
	data.Artists = artists

	app.render(w, http.StatusOK, "create.tmpl", data)

}

// viewRecord - displays single record
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

	for _, r := range records {
		artist, _ := app.artists.GetRecordArtist(int(r.ID))
		r.RecordArtist = append(r.RecordArtist, artist...)
	}

	data := app.newTemplateData(r)
	data.Records = records

	app.render(w, http.StatusOK, "records.tmpl", data)

}
