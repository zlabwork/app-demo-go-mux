package middleware

import (
	"app"
	"app/libs/utils"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"time"
)

var logger = logrus.New()

func init() {

	logger.SetOutput(&lumberjack.Logger{
		Filename:   app.Dir.Data + "access.log",
		MaxSize:    200,
		MaxBackups: 20,
		MaxAge:     180,
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

		id := r.Header.Get("Trace-Id")
		if id == "" {
			id = utils.NewObjectID().Hex()
			w.Header().Set("Trace-Id", id)
		}

		date := time.Now().Format(time.RFC3339)
		logger.Println(fmt.Sprintf("[%s] [%s] %s \"%s %s\" \"%s\"", date, id, getRealIP(r), r.Method, r.RequestURI, r.Header.Get("User-Agent")))

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		ctx := context.WithValue(r.Context(), app.TraceKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
