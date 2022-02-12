package main

import (
	"app"
	"app/middleware"
	"app/restful"
	"context"
	"flag"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	// app.yaml
	bs, err := ioutil.ReadFile("../config/app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if yaml.Unmarshal(bs, app.Yaml) != nil {
		log.Fatal(err)
	}

	// libs
	app.Libs = app.NewLibs()

	// config & router
	var wait time.Duration
	var dir string
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&dir, "dir", "assets", "the directory to serve files")
	flag.Parse()

	r := mux.NewRouter().StrictSlash(true)
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.SignatureMiddleware)
	r.HandleFunc("/", restful.DefaultHandler)
	r.HandleFunc("/demo1/{method:[a-z]+}/{name:[0-9a-zA-Z_-]+}", restful.DefaultHandler).Methods(http.MethodGet)
	r.HandleFunc("/demo2/{id:[0-9]+}", restful.DefaultHandler).Methods(http.MethodGet, http.MethodPut)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(dir))))
	r.NotFoundHandler = http.HandlerFunc(restful.NotFoundHandler)

	srv := &http.Server{
		Handler:      r,
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
	app.Banner("Service port :" + os.Getenv("APP_PORT"))
	log.Println("the service is started")

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

func init() {
	f, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(f)
	}
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})
}
