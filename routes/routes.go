package routes

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/cpustejovsky/catchall/handlers"
	"github.com/cpustejovsky/catchall/logger"
	"github.com/cpustejovsky/catchall/middleware"
	"github.com/justinas/alice"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(log logger.Logger, client *mongo.Client) http.Handler {

	middlewares := middleware.Middleware{
		Logger: log,
	}

	standardMiddleware := alice.New(middlewares.RecoverPanic, middlewares.LogRequest, middlewares.SecureHeaders)

	mux := pat.New()

	database := client.Database("catchall_domains")
	collection := database.Collection("domains")

	domainHandlers := handlers.Handler{
		Logger:     log,
		Collection: collection,
	}
	mux.Put("/events/:domain_name/delivered", standardMiddleware.ThenFunc(domainHandlers.UpdateDelivered))
	mux.Put("/events/:domain_name/bounced", standardMiddleware.ThenFunc(domainHandlers.UpdateBounced))
	mux.Get("/domains/:domain_name", standardMiddleware.ThenFunc(domainHandlers.CheckStatus))
	mux.Get("/ping", standardMiddleware.ThenFunc(domainHandlers.Ping))

	return standardMiddleware.Then(mux)
}
