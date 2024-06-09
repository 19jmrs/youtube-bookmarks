package main

import (
	"fmt"
	"net/http"

	"github.com/19jmrs/youtube-bookmarks/views"
	"github.com/a-h/templ"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("hello from main")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fileServer))

	r.Get("/", templ.Handler(views.Landing()).ServeHTTP)
	

	http.ListenAndServe(":8000", r)

}
