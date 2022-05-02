package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("dist"))

	mux.Handle("/dist/", http.StripPrefix("/dist", fileServer))

	mux.HandleFunc("/admin", homeAdmin)
	mux.HandleFunc("/admin/records", viewRecords)
	mux.HandleFunc("/admin/records/view", viewRecord)
	mux.HandleFunc("/admin/record/create", createRecordPost)
	mux.HandleFunc("/admin/record/new", createRecordGet)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalln(err)
	}

}
