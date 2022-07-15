package middleware

import (
	"app"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
)

var logger = logrus.New()

func init() {

	logger.SetOutput(&lumberjack.Logger{
		Filename:   app.Dir.Data + "access.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := r.Header.Get("Request-Id")
		if id == "" {
			id = uuid.New().String()
			w.Header().Set("Request-Id", id)
		}

		// Do stuff here
		logger.Println(r.RequestURI)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
