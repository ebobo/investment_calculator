package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func (s *Server) startHTTP() error {
	m := mux.NewRouter()

	// Add CORS
	cors := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		MaxAge:           31,
		Debug:            false,
	})

	// Add the REST machinery for grpc-gateway into the Gorilla mux.
	restMux, err := s.loanService.RESTMuxViaGRPC(s.ctx, s.grpcListenAddr)
	if err != nil {
		return err
	}

	// This is where you add other stuff you want to map in the mux

	// Add handler for REST interface
	m.PathPrefix("/api/v1").Handler(restMux)

	httpServer := &http.Server{
		Addr:              s.httpListenAddr,
		Handler:           handlers.ProxyHeaders(cors.Handler(m)),
		ReadTimeout:       (10 * time.Second),
		ReadHeaderTimeout: (8 * time.Second),
		WriteTimeout:      (45 * time.Second),
	}

	// Set up shutdown handler
	go func() {
		<-s.ctx.Done()
		err := s.loanService.CloseInternalClientConn()
		if err != nil {
			log.Printf("error closing internal grpc client: %v", err)
		}

		err = httpServer.Shutdown(context.Background())
		if err != nil {
			log.Printf("error shutting down HTTP interface '%s': %v", s.httpListenAddr, err)
		}
	}()

	// Start HTTP server
	go func() {
		log.Printf("starting HTTP interface '%s'", s.httpListenAddr)

		// This isn't entirely true and really represents a race condition, but
		// doing this properly is a pain in the neck.
		s.httpStarted.Done()

		err := httpServer.ListenAndServe()
		if err == http.ErrServerClosed {
			err = errors.New("")
		}

		log.Printf("HTTP interface '%s' down %v", s.httpListenAddr, err)
		s.httpStopped.Done()
	}()

	return nil
}
