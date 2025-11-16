#!/bin/bash

# Script to start the Go application with Air (hot reload)
air --build.cmd "go build -o ./bin/api ./cmd/api" --build.bin "./bin/api"
