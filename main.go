package main

import (
	"go.uber.org/zap"
	"kvdb/storage/engine"
	"log"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	db := engine.New(logger)
	db.Cli()
}
