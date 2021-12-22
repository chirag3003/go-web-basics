package main

import (
	"github.com/bmizerany/pat"
	"github.com/chirag3003/go-web/pkg/config"
	"github.com/chirag3003/go-web/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}