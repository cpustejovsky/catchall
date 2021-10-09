package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()

	mux.Put("/events/:domain_name/delivered", standardMiddleware.ThenFunc(app.updateDelivered))
	mux.Put("/events/:domain_name/bounced", standardMiddleware.ThenFunc(app.updateBounced))
	mux.Get("/domains/:domain_name", standardMiddleware.ThenFunc(app.checkStatus))

	return standardMiddleware.Then(mux)
}
