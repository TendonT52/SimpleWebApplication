package main

import (
	"log"
	"net/http"

	"github.com/TendonT52/SimpleWebApplication/pkg/config"
	"github.com/TendonT52/SimpleWebApplication/pkg/handlers"
	"github.com/TendonT52/SimpleWebApplication/pkg/render"
)

const portNumber string = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
