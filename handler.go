package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bmanth60/DuckTracker/api"
	"github.com/bmanth60/DuckTracker/data"
	dterr "github.com/bmanth60/DuckTracker/errors"
	"github.com/bmanth60/DuckTracker/types"
)

//Handler http routing handler
type Handler struct {
	Db  *data.Database
	API *api.API
}

//ServeHTTP serve the http request
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch {
	case r.URL.Path == "/":
		err = h.HandleIndex(w, r)
		break
	default:
		err = dterr.ErrPageNotFound
	}

	if err != nil && err == dterr.ErrPageNotFound {
		http.NotFoundHandler().ServeHTTP(w, r)
	} else if err != nil {
		log.Println(err)
		http.Error(w, "unable to complete request", http.StatusInternalServerError)
		return
	}
}

//HandleIndex handle index/root path requests
func (h *Handler) HandleIndex(w http.ResponseWriter, r *http.Request) error {
	p := &types.Page{Title: "Duck Tracker"}

	// Check if this is a postback
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			// TODO - Brian - 20181009 - Handle this error better
			return err
		}

		p.Messages = h.API.HandleForm(r.PostForm, h.Db.AddDuckEntry)
	} else if r.Method != http.MethodGet {
		return dterr.ErrPageNotFound
	}

	// Run template
	t, err := template.ParseFiles("presentation/index.html")
	if err != nil {
		return err
	}
	t.Execute(w, p)

	return nil
}
