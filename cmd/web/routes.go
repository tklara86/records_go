package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("dist"))

	router.Handler(http.MethodGet, "/dist/*filepath", http.StripPrefix("/dist", fileServer))

	router.HandlerFunc(http.MethodGet, "/admin", app.homeAdmin)
	router.HandlerFunc(http.MethodGet, "/admin/records", app.viewRecords)
	router.HandlerFunc(http.MethodGet, "/admin/records/view/:id", app.viewRecord)
	router.HandlerFunc(http.MethodPost, "/admin/record/new", app.recordCreatePost)
	router.HandlerFunc(http.MethodGet, "/admin/record/new", app.recordCreateGet)
	//JSON
	router.HandlerFunc(http.MethodGet, "/admin/record/labelsJSON", app.getLabelsJSON)
	router.HandlerFunc(http.MethodGet, "/admin/record/artistsJSON", app.getArtistsJSON)
	router.HandlerFunc(http.MethodGet, "/admin/record/genresJSON", app.getGenresJSON)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)

}
