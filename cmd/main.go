package main

import (
	"context"
	"golang-api/ent"
	"golang-api/internal/config"
	"golang-api/internal/registry"
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var c config.Config
	cleanenv.ReadEnv(&c)

	client, err := ent.Open(config.DatastoreType, c.Datastore.File)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	server, err := registry.InitializeServer(c, client)
	if err != nil {
		log.Fatalf("failed initialize server: %v", err)
	}
	err = server.Start()
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
