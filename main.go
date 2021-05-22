package main

import (
	"go-api-dev/app"
	"go-api-dev/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
