package main

import (
	"log"
	"net/http"
	"time"

	"github.com/atomicmeganerd/rcd-gopher-social/internal/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type config struct {
	addr string
	db   dbConfig
	env  string
}

type application struct {
	config config
	store  store.Storage
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdeConns  int
	maxIdleTime  string
}

// chi.Router implements the http.Handler interface
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// This middleware sets a request ID for each request, which we can use to trace the request
	r.Use(middleware.RequestID)

	// This middleware sets the X-Real-IP header to the client's real IP address, if it's present
	r.Use(middleware.RealIP)

	// This middleware logs the start and end of each request, how long it took, and the HTTP
	// status code
	r.Use(middleware.Logger)

	// This middleware recovers from panics, logs the panic, and returns a 500 Internal Server Error
	r.Use(middleware.Recoverer)

	// This middleware sets a timeout value on the request context, and automatically responds with
	// a 504 Gateway Timeout if the handler takes too long to complete
	r.Use(middleware.Timeout(60 * time.Second))

	// We can group routes by version
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Println("Starting server on", app.config.addr)
	return srv.ListenAndServe()
}
