package internal

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Run() {
	r := chi.NewRouter()

	r.Get("/", Index)

	http.ListenAndServe(":8080", r)
}
