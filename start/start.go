package main

import (
	"adm-num/ent"
	_ "adm-num/ent/runtime"
	"adm-num/routes"
	"adm-num/utils"
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/labstack/echo/v4"
)

func main() {

	ctx := context.Background()

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	// Create an ent.Client with Postgres database.
	client, err := ent.Open(dialect.Postgres, config.DbSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgress: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := echo.New()

	routes.Init(app, client, config)

	app.Logger.Fatal(app.Start(config.ServerAddr))

}
