package handler

import (
	"net/http"

	"github.com/thenopholo/go-bid/internal/config"
)

type Handler struct{
  logger *config.Logger
}

func NewHandlrer() *Handler  {
  logger := config.NewLogger("HANDLER")
  return &Handler{
  	logger: logger,
  }
}


func HealthCheck(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("OK"))
}