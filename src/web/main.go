package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/yasniel1408/bookings/src/config"
	"github.com/yasniel1408/bookings/src/controllers"
	"github.com/yasniel1408/bookings/src/models"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Staring application on port %s", portNumber)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	server.ListenAndServe()
}

func run() error {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	//change this to true when in produccion
	app.InProduccion = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour // session de 24 horas
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false // en produccion deberia ser true, esto es ahora para conextarnos a localhost:8080

	app.Session = session

	tc, err := config.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := controllers.NewRepo(&app)
	controllers.NewHandlers(repo)

	config.NewTemplates(&app)

	return nil
}
