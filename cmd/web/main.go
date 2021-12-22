package main

import (
	"fmt"
	"github.com/chirag3003/go-web/pkg/config"
	"github.com/chirag3003/go-web/pkg/handlers"
	"github.com/chirag3003/go-web/pkg/render"
	"log"
	"net/http"
)

const PORT = ":3000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}

}
