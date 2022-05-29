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

	// RecordArtists
	_, err = app.artists.InsertRecordArtist(recordArtist)
	if err != nil {
		app.serverError(w, err)
		return
	}

	recordLabel := []models.RecordLabel{}
	for _, l := range r.PostForm["label-name"] {
		labelID, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		rl := []models.RecordLabel{
			{
				LabelID:  labelID,
				RecordID: int64(id),
			},
		}

		recordLabel = append(recordLabel, rl...)
	}

	// RecordLabels
	_, err = app.labels.InsertRecordLabel(recordLabel)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// http.Redirect(w, r, "/admin/records", http.StatusFound)

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

	// pass labels struct
	labels, err := app.labels.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)

	data.Genres = genres
	data.Artists = artists
	data.Labels = labels

	app.render(w, http.StatusOK, "create.tmpl", data)

}

func (app *application) getLabelsJSON(w http.ResponseWriter, r *http.Request) {

	labels, err := app.labels.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, l := range labels {
		l.InputName = "label-name"
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"labels": labels}, nil)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) postLabelsJSON(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name string `json:"name"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, err)
	}

	label := &models.Label{
		Name: input.Name,
	}

	_, err = app.labels.Insert(label)
	if err != nil {
		app.serverError(w, err)
	}

	headers := make(http.Header)

	err = app.writeJSON(w, http.StatusCreated, envelope{"label": label}, headers)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) postArtistsJSON(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name string `json:"name"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, err)
	}

	artist := &models.Artist{
		Name: input.Name,
	}

	_, err = app.artists.Insert(artist)
	if err != nil {
		app.serverError(w, err)
	}

	headers := make(http.Header)

	err = app.writeJSON(w, http.StatusCreated, envelope{"artist": artist}, headers)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) postGenresJSON(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name string `json:"name"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverError(w, err)
	}

	genre := &models.Genre{
		GenreName: input.Name,
	}

	_, err = app.genres.Insert(genre)
	if err != nil {
		app.serverError(w, err)
	}

	headers := make(http.Header)

	err = app.writeJSON(w, http.StatusCreated, envelope{"genre": genre}, headers)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) getArtistsJSON(w http.ResponseWriter, r *http.Request) {

	artists, err := app.artists.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, a := range artists {
		a.InputName = "artist-name"
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"artists": artists}, nil)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) getGenresJSON(w http.ResponseWriter, r *http.Request) {

	genres, err := app.genres.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, g := range genres {
		g.InputName = "genre-name"
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"genres": genres}, nil)
	if err != nil {
		app.serverError(w, err)
	}

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

// viewRecords - displays all records
func (app *application) viewRecords(w http.ResponseWriter, r *http.Request) {

	records, err := app.records.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// attach artist(s), labels(s) to record
	for _, r := range records {
		artist, _ := app.artists.GetRecordArtist(int(r.ID))
		label, _ := app.labels.GetLabelArtist(int(r.ID))

		r.RecordArtist = append(r.RecordArtist, artist...)
		r.RecordLabel = append(r.RecordLabel, label...)
	}

	data := app.newTemplateData(r)
	data.Records = records

	app.render(w, http.StatusOK, "records.tmpl", data)

}
