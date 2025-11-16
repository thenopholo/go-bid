package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/thenopholo/go-bid/internal/config"
	"github.com/thenopholo/go-bid/internal/handler"
)

type Server struct {
	port    string
	mux     *chi.Mux
	handler *handler.Handler
	server  *http.Server
	logger  *config.Logger
}

func NewServer() *Server {
	logger := config.NewLogger("SERVER")
	r := chi.NewRouter()
	h := handler.NewHandlrer()
	port := os.Getenv("SERVER_PORT")

	s := &http.Server{
		Addr:              ":" + port,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return &Server{
		port:    port,
		mux:     r,
		handler: h,
		server:  s,
		logger:  logger,
	}
}

func (s *Server) SetupRoutes() {
	Router(s.mux, s.handler)
}

func (s *Server) Start() {
	s.SetupRoutes()
	s.server.Handler = s.mux

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		s.logger.Info("Server UP")
		if err := s.server.ListenAndServe(); err != nil {
			s.logger.Errf("Server failed: %v", err.Error())
		}
	}()

	<-stop
	s.Shutdown()
}

func (s *Server) Shutdown() {
	s.logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Errf("Server forced to shutdown: %v", err.Error())
	}

	s.logger.Info("Server stopped gracefully")
}
