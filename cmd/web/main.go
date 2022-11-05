package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ShohMansurjonovich/bookings/pkg/config"
	"github.com/ShohMansurjonovich/bookings/pkg/handler"
	"github.com/ShohMansurjonovich/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)
	fmt.Println("Running port: ", portNumber)
	// http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal("Error occured in server!")
	}

}
