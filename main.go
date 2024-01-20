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
	r.Get("/employees",handlers.H2)
	r.Get("/employee-details",handlers.H3)
	r.Get("/employee-details",handlers.H4)
	r.Get("/employee-technical",handlers.H5)
	r.Get("/employee-documents",handlers.H6)
	r.Get("/employee-salary",handlers.H7)
	r.Get("/transact-salaries",handlers.H8)
	

	log.Println("Listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
