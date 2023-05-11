package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/render"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig //added here to meet with the scope of the struct globally.
var session *scs.SessionManager

// main is the main function
func main() {

	//chage var to true when in production
	app.InProduction = false
	//session packages to set time limit to user login
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true                  //makes sure to persist during user window acitvity
	session.Cookie.SameSite = http.SameSiteLaxMode //how strict should mode be to log user out
	session.Cookie.Secure = app.InProduction       //false for now, but true in production to ensre secutity in https connection

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot do")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
