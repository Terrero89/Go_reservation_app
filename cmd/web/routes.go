package main

import (
	"github.com/Terrero89/reservations_app/pkg/config"
	"github.com/Terrero89/reservations_app/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//creating a file server to get images and other docs
	fileServer := http.FileServer(http.Dir("./static/images")) //created file with where and name of the directory

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer)) // select all object or files within the static a
	return mux
}
