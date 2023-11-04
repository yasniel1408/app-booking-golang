package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yasniel1408/bookings/pkg/config"
	"github.com/yasniel1408/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//Middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// FileServer
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
