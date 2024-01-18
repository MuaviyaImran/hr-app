package main

import (
	"hr/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)


func main() {
	r := chi.NewRouter()

	r.Handle("/css/*", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
    r.Get("/",handlers.H1)

	log.Println("Listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
