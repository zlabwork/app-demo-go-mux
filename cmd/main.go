package main

import (
	"app"
	"app/middleware"
	"app/restful"
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {

	// rand seed
	rand.Seed(time.Now().UnixNano())

	// env
	err := godotenv.Load(app.Dir.Root + ".env")
	if err != nil {
		log.Fatal(err)
	}

	// app.yaml
	bs, err := ioutil.ReadFile(app.Dir.Config + "app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if yaml.Unmarshal(bs, app.Yaml) != nil {
		log.Fatal(err)
	}

	// logs
	f, err := os.OpenFile(app.Dir.Data+"system.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(f)
	}
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})

	// libs
	app.Libs = app.NewLibs()
}

func usage() {

	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func router() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.SignatureMiddleware)
	r.HandleFunc("/", restful.DefaultHandler)
	r.HandleFunc("/demo1/{method:[a-z]+}/{name:[0-9a-zA-Z_-]+}", restful.DefaultHandler).Methods(http.MethodGet)
	r.HandleFunc("/demo2/{id:[0-9]+}", restful.DefaultHandler).Methods(http.MethodGet, http.MethodPut)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(app.Dir.Assets))))
	r.NotFoundHandler = http.HandlerFunc(restful.NotFoundHandler)
	// router group
	user := r.PathPrefix("/users").Subrouter()
	user.HandleFunc("/{id:[0-9a-zA-Z_-]+}", restful.DefaultHandler).Methods(http.MethodGet)
	return r
}

func main() {

	var help bool
	var wait time.Duration
	flag.Usage = usage
	flag.BoolVar(&help, "h", false, "help")
	flag.DurationVar(&wait, "timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	if help {
		usage()
		os.Exit(1)
	}

	srv := &http.Server{
		Handler:      router(),
		Addr:         "0.0.0.0:" + os.Getenv("APP_PORT"),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	app.Banner("Listening On :" + os.Getenv("APP_PORT"))
	log.Println("service is started")

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
