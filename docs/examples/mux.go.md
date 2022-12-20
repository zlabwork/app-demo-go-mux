```go
package main

import (
    "app"
    "context"
    "flag"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
)

func main() {
    // .env
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal(err)
    }
    // app.yaml
    bs, err := ioutil.ReadFile("config/app.yaml")
    err = yaml.Unmarshal(bs, app.Yaml)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    // config & router
    var wait time.Duration
    var dir string
    flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
    flag.StringVar(&dir, "dir", "assets", "the directory to serve files")
    flag.Parse()

    r := mux.NewRouter()
    r.HandleFunc("/", DefaultHandler)
    r.HandleFunc("/events/{method:[a-z]+}/{name:[0-9a-zA-Z_-]+}", web.EventsHandler).Methods("GET")
    r.HandleFunc("/publish/{id:[0-9]+}", web.PublishHandler).Methods("GET")
    r.HandleFunc("/project", web.ProjectHandler).Methods("GET", "POST")
    r.HandleFunc("/project/{id:[0-9]+}", web.ProjectModifyHandler).Methods("GET", "PUT")
    r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(dir))))

    srv := &http.Server{
        Handler:      r,
        Addr:         "0.0.0.0:8000",
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

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
}
```
