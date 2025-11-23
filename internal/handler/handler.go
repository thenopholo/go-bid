package handler

import (
	"net/http"

	"github.com/thenopholo/go-bid/internal/config"
	"github.com/thenopholo/go-bid/internal/service"
)

type Handler struct {
	UserService *service.UserService
	logger      *config.Logger
}

func NewHandlrer(userService *service.UserService) *Handler {
	logger := config.NewLogger("HANDLER")
	return &Handler{
    UserService: userService,
		logger: logger,
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}