package middleware

import (
	"app"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"time"
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
	logger.SetFormatter(new(logsFormatter))
}

// custom logs formatter
type logsFormatter struct{}

func (f *logsFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message + "\n"), nil
}

func getRealIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarder-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := r.Header.Get("Request-Id")
		if id == "" {
			id = uuid.New().String()
			w.Header().Set("Request-Id", id)
		}

		dt := time.Now().Format(time.RFC3339)
		logger.Println(fmt.Sprintf("%s [%s] \"%s %s\" \"%s\"", getRealIP(r), dt, r.Method, r.RequestURI, r.Header.Get("User-Agent")))

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
