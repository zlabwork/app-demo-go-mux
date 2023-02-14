package middleware

import (
	"app"
	"app/libs/utils"
	"app/msg"
	"app/response"
	"context"
	"net/http"
	"strings"
)

const (
	headerKeyToken = "X-Lab-Token"
)

func AuthenticateMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// ignore
		prefix := []string{"/keys", "/token", "/assets"}
		for _, p := range prefix {
			if strings.HasPrefix(r.URL.Path, p) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// parse
		token := r.Header.Get(headerKeyToken)
		if token == "" {
			response.Message(r.Context(), w, msg.ErrAccess, "missing "+headerKeyToken)
			return
		}
		claims, err := utils.ParseJWT(token)
		if err != nil {
			response.Message(r.Context(), w, msg.ErrAccess, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), app.AuthKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
