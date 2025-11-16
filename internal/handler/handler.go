package handler

import "net/http"

type Handler struct{}

func NewHandlre() *Handler {
  return &Handler{}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("OK"))
}