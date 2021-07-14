package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Start() {
	r := chi.NewRouter()

	http.ListenAndServe("", r)
}
