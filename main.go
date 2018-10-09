package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/bmanth60/DuckTracker/types"
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

	log.Println("Starting server...")
	//services := new(Services)

	http.HandleFunc("/", ServeHTTP)

	setup := getSetupChannel()
	if setup != nil {
		setup <- "done"
	}

	log.Println("Listening on port 80...")
	log.Fatal(http.ListenAndServe(":80", logRequest(http.DefaultServeMux)))
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := &types.Page{Title: "Hello World"}

	// Run template
	t, err := template.ParseFiles("presentation/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, p)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL))
		handler.ServeHTTP(w, r)
	})
}
