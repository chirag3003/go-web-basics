package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/chirag3003/go-web/pkg/config"
	"github.com/chirag3003/go-web/pkg/handlers"
	"github.com/chirag3003/go-web/pkg/render"
	"log"
	"net/http"
	"time"
)

const PORT = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false //change it to true when in production

	//session configs
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	//html template config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplate(&app)

	//HTTP Handlers config
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	//err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}
	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}
	log.Fatal(srv.ListenAndServe())

}
