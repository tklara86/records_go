package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}
	// Initialize a new buffer.
	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)

}

type Link struct {
	LinkTitle string
	LinkPath  string
}

type SideBarLink struct {
	LinkTitle string
	LinkIcon  string
	LinkHref  string
}

type FormOption struct {
	FormName string
	FormType string
	FormFor  string
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
		Links: []Link{
			{
				LinkTitle: "Home",
				LinkPath:  "/admin",
			},
			{
				LinkTitle: "Categories",
				LinkPath:  "/admin/categories/",
			},
		},
		SideBarLinks: []SideBarLink{
			{
				LinkTitle: "Home",
				LinkIcon:  "home",
				LinkHref:  "/admin",
			},
			{
				LinkTitle: "Albums",
				LinkIcon:  "archive",
				LinkHref:  "/admin/records",
			},
		},
		FormOptions: []FormOption{
			{
				FormName: "Title",
				FormType: "text",
				FormFor:  "recordTitle",
			},
			{
				FormName: "Release Date",
				FormType: "text",
				FormFor:  "recordReleaseDate",
			},
			// {
			// 	FormName: "Label",
			// 	FormType: "text",
			// 	FormFor:  "recordLabel",
			// },
		},
	}
}

type envelope map[string]interface{}

func (app *application) writeJSON(w http.ResponseWriter, status int,
	data envelope, headers http.Header) error {

	json, err := json.MarshalIndent(data, " ", "   ")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	_, err = w.Write(json)
	if err != nil {
		return err
	}

	return nil

}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// limit the size of the request body to 1MB
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %s)", syntaxError)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err

		}

	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}
	return nil
}
