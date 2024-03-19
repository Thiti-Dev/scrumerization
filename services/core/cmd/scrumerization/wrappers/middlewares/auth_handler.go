package middlewares

import (
	"context"
	"net/http"

	context_type "github.com/Thiti-Dev/scrumerization-core-service/internal/domain/context"
)

func AuthHandlerForGraphql(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		// if auth is not available then proceed to resolver
		if header == "" {
			next.ServeHTTP(w, r)
		} else {
			ctx := context.WithValue(r.Context(), context_type.TokenCtxKey, header) // just forward it, will do the check in the directive
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
