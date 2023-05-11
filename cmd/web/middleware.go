package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit next page")
		next.ServeHTTP(w, r)
	})
}

// protect to post request CORS
func NoSsuf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// middleware to load and saves session in every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next) // this line automatically load and saves session information for current req in a cookie
}
