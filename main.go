package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bmanth60/DuckTracker/api"
	"github.com/bmanth60/DuckTracker/data"
)

var (
	gStatus chan string
)

func setSetupChannel(status chan string) chan string {
	gStatus = status
	return gStatus
}

func getSetupChannel() chan string {
	return gStatus
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(fmt.Sprintf("Project %s:%s", os.Getenv("PROJECT_NAME"), os.Getenv("PROJECT_BUILD")))

	log.Println("Starting database...")
	db, err := data.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server...")
	handler := new(Handler)
	handler.Db = db
	handler.API = new(api.API)

	http.Handle("/", handler)

	setup := getSetupChannel()
	if setup != nil {
		setup <- "done"
	}

	log.Println("Listening on port 80...")
	log.Fatal(http.ListenAndServe(":80", logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL))
		handler.ServeHTTP(w, r)
	})
}
