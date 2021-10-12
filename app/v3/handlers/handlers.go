package handlers

import (
	"context"
	"expvar"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/ardanlabs/service/business/web/v1/mid"
	"github.com/ardanlabs/service/foundation/web"
	"github.com/cpustejovsky/catchall/app/v3/handlers/debug/checkgrp"
	"github.com/cpustejovsky/catchall/app/v3/handlers/domaingrp"
	"github.com/cpustejovsky/catchall/internal/core/domain/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Options represent optional parameters.
type Options struct {
	corsOrigin string
}

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	Client   *mongo.Client
}

type RoutesConfig struct {
	Log    *zap.SugaredLogger
	Client *mongo.Client
}

// Routes binds all the version 1 routes.
func Routes(app *web.App, cfg RoutesConfig) {
	const version = "v3"

	// Register user management and authentication endpoints.
	domainCollection := cfg.Client.Database("catchall_domainsv3").Collection("domainsv3")
	store := db.NewStore(cfg.Log, domainCollection)
	dgh := domaingrp.Handlers{
		Store: store,
	}

	app.Handle(http.MethodPut, version, "/events/:domain_name/delivered", dgh.UpdateBounced)
	app.Handle(http.MethodPut, version, "/events/:domain_name/bounced", dgh.UpdateDelivered)
	app.Handle(http.MethodGet, version, "/domains/:domain_name", dgh.CheckStatus)
	app.Handle(http.MethodGet, version, "/ping", Ping)
}

func Ping(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Hello, World")
	w.Write([]byte("OK"))
	return nil
}

// APIMux constructs an http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig, options ...func(opts *Options)) http.Handler {
	var opts Options
	for _, option := range options {
		option(&opts)
	}

	// Construct the web.App which holds all routes as well as common Middleware.
	app := web.NewApp(
		cfg.Shutdown,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Metrics(),
		mid.Panics(),
	)

	// Accept CORS 'OPTIONS' preflight requests if config has been provided.
	// Don't forget to apply the CORS middleware to the routes that need it.
	// Example Config: `conf:"default:https://MY_DOMAIN.COM"`
	if opts.corsOrigin != "" {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			return nil
		}
		app.Handle(http.MethodOptions, "", "/*", h)
	}

	Routes(app, RoutesConfig{
		Log:    cfg.Log,
		Client: cfg.Client,
	})

	return app
}

// DebugStandardLibraryMux registers all the debug routes from the standard library
// into a new mux bypassing the use of the DefaultServerMux. Using the
// DefaultServerMux would be a security risk since a dependency could inject a
// handler into our service without us knowing it.
func DebugStandardLibraryMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Register all the standard library debug endpoints.
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())

	return mux
}

// DebugMux registers all the debug standard library routes and then custom
// debug application routes for the service. This bypassing the use of the
// DefaultServerMux. Using the DefaultServerMux would be a security risk since
// a dependency could inject a handler into our service without us knowing it.
func DebugMux(build string, log *zap.SugaredLogger, client *mongo.Client) http.Handler {
	mux := DebugStandardLibraryMux()

	// Register debug check endpoints.
	cgh := checkgrp.Handlers{
		Build:  build,
		Log:    log,
		Client: client,
	}
	mux.HandleFunc("/debug/readiness", cgh.Readiness)
	mux.HandleFunc("/debug/liveness", cgh.Liveness)

	return mux
}
