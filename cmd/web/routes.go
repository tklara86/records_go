package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("dist"))

	mux.Handle("/dist/", http.StripPrefix("/dist", fileServer))

	mux.HandleFunc("/admin", app.homeAdmin)
	mux.HandleFunc("/admin/records", app.viewRecords)
	mux.HandleFunc("/admin/records/view", app.viewRecord)
	mux.HandleFunc("/admin/record/create", app.createRecordPost)
	mux.HandleFunc("/admin/record/new", app.createRecordGet)

	return mux
}
