package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/sebenitezg/hotel-service/config"
)

// HTTPServer http server
type HTTPServer struct {
	sc     config.ServerConfigurations
	Router *chi.Mux
}

func NewHTTPServer(serverConf config.ServerConfigurations) *HTTPServer {
	router := chi.NewRouter()

	// APM middleware
	//router.Use(apmchiv5.Middleware())

	// A good base middleware stack
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request models (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	return &HTTPServer{
		sc:     serverConf,
		Router: router,
	}
}

func (r *HTTPServer) Start() {
	listeningAddr := ":" + r.sc.Port
	log.Printf("Server listening on port %s", listeningAddr)

	// Customizing the server
	s := &http.Server{
		Addr:         listeningAddr,
		Handler:      r.Router,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	// Start the server
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start http server. %v", err)
	}
}
