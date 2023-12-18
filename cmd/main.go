package main

import (
	"context"
	"golang-api/ent"
	"golang-api/internal/config"
	"golang-api/internal/registry"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	var c config.Config
	cleanenv.ReadEnv(&c)

	client, err := ent.Open(config.DatastoreType, c.Datastore.File)
	if err != nil {
		logger.Fatal("failed opening connection to sqlite", zap.Error(err))
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal("failed creating schema resources", zap.Error(err))
	}

	server, err := registry.InitializeServer(c, client, logger)
	if err != nil {
		logger.Fatal("failed initialize server", zap.Error(err))
	}
	err = server.Start()
	if err != nil {
		logger.Fatal("failed to run server", zap.Error(err))
	}
}
