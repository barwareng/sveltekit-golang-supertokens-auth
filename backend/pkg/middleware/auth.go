package middleware

import (
	"net/http"

	"github.com/supertokens/supertokens-golang/recipe/session"
)

func VerifySession(handler http.Handler) http.Handler {
	return session.VerifySession(nil, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))
}
