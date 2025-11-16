package main

import (
	"github.com/thenopholo/go-bid/internal/config"
	"github.com/thenopholo/go-bid/internal/server"
)

func main() {
	logger := config.NewLogger("MAIN_API")
	logger.Info("Starting application...")
	InitServer()
}

func InitServer() {
	server.NewServer().Start()
}
