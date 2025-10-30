package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
	"github.com/thenopholo/go-bid/internal/config"
)

func main() {
	logger := config.NewLogger("MAIN_TERN")

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/migrations",
		"--config",
		"./internal/store/migrations/tern.conf",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Errf("erro on migrate: %v", err)
		logger.Debugf("recived output: %s", string(output))
		panic(err)
	}

	fmt.Printf("Successfully migrate output: %s", string(output))
}
