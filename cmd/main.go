package main

import (
	"app"
	"app/libs/utils"
	"app/middleware"
	"app/restful"
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func usage() {

	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func router() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	// Logs
	r.Use(middleware.LoggingMiddleware)

	// Guard
	if os.Getenv("AUTH_GUARD") == "token" {
		r.Use(middleware.AuthenticateMiddleware)
	} else if os.Getenv("AUTH_GUARD") == "sign" {
		r.Use(middleware.SignatureMiddleware)
	}

	r.HandleFunc("/", restful.DefaultHandler)
	r.HandleFunc("/demo1/{method:[a-z]+}/{name:[0-9a-zA-Z_-]+}", restful.DefaultHandler)
	r.HandleFunc("/demo2/{id:[0-9]+}", restful.DefaultHandler)
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
	ips, _ := utils.LocalIPv4s()
	output := `
http://localhost:%s
http://%s:%s`
	app.Banner(fmt.Sprintf(output, os.Getenv("APP_PORT"), ips[0], os.Getenv("APP_PORT")))
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
