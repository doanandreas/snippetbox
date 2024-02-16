package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	r.Use(app.recoverPanic, app.logRequest, secureHeaders)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))

	r.Get("/", app.home)
	r.Get("/snippet/view/{id}", app.snippetView)
	r.Get("/snippet/create", app.snippetCreate)
	r.Post("/snippet/create", app.snippetCreatePost)

	return r
}
