package mw

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	validCookie := true
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging:", r.Method, r.URL.Path)
		_, err := r.Cookie("UserID")
		if err != nil {
			validCookie = false
		}

		if validCookie {
			next.ServeHTTP(w, r)
		} else {
			// redirect to login
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	})
}
