package main

import "GophKeeper/internal/logger"

func main() {
	log := logger.NewLogger()
	log.Info("starting server")
}
