package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func h1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	message := "Authenticated successfully!"
	fmt.Fprintf(w, message)
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static files
	r.Handle("/css/*", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Routes
	r.Get("/", h1)
	r.Get("/api/auth", handleAuth)

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
