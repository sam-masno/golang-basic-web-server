package main

import (
	"basic-server/pkg/config"
	"basic-server/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(NoSurf)
	r.Use(sessions)
	r.Use(sessionRemoteAddr)
	r.Use(writeToConsole)
	r.Get("/", handlers.Repo.HomePage)
	r.Get("/about", handlers.Repo.About)

	return r
}
