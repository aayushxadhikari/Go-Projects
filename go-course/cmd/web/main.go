package main

// packages always live in their own directory
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aayushxadhikari/go-course/pkg/config"
	"github.com/aayushxadhikari/go-course/pkg/handlers"
	"github.com/aayushxadhikari/go-course/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8081"
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/hello", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("%s", fmt.Sprintf("Starting application on port %s", portNumber))
	// http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr: portNumber,
		Handler:routes(&app),

	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}