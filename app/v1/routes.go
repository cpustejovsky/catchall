package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()

	mux.Put("/events/:domain_name/delivered", http.HandlerFunc(app.updateDelivered))
	mux.Put("/events/:domain_name/bounced", http.HandlerFunc(app.updateBounced))
	mux.Get("/domains/:domain_name", http.HandlerFunc(app.checkStatus))
	mux.Get("/ping", http.HandlerFunc(app.ping))

	return mux
}
