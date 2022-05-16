package main

import (
	logger "prometheus-http-sd/core"
	"prometheus-http-sd/handlers"
)

func main() {

	logger, dispose := logger.New()
	defer dispose()

	logger.Info("Starting prometheus-http-sd")

	handlers.HttpHandler(logger)
}
