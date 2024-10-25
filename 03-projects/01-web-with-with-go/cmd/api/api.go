package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gcalvocr/go-stripe/internal/driver"
)

const version = "1.0.0"

// configuration of the application
type config struct {
	port int
	env  string
	db   struct {
		dsn string // data source name
	}
	stripe struct {
		secret string
		key    string
	}
}

// The receiver for the various part of the application
type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting Backend Server in %s on port %d\n", app.config.env, app.config.port)

	return srv.ListenAndServe()
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "dev", "Application environment {dev|prod|stage}")
	flag.StringVar(&cfg.db.dsn, "dns", "gabriel:secret@tcp(localhost:3306)/widgets?parseTime=true&tls=false", "DSN")
	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Connection to DB
	conn, err := driver.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
