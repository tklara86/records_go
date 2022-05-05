package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":8000", "HTTP Netword Address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)

	errorLog := log.New(os.Stderr, "Error\t", log.LstdFlags)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("dist"))

	mux.Handle("/dist/", http.StripPrefix("/dist", fileServer))

	mux.HandleFunc("/admin", homeAdmin)
	mux.HandleFunc("/admin/records", viewRecords)
	mux.HandleFunc("/admin/records/view", viewRecord)
	mux.HandleFunc("/admin/record/create", createRecordPost)
	mux.HandleFunc("/admin/record/new", createRecordGet)

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Server starting on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
