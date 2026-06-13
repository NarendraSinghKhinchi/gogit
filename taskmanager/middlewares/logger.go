package middlewares

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	handler := func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		fmt.Println(r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
}
