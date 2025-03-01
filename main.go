package main

import (
	"backend/internal/core"
	"backend/internal/router"
	"backend/pkg/application"
	"log"
)

// @title Pokedex
// @version 1.0.0
// @description Un Pokédex en golang
// @BasePath /
func main() {
	app := application.NewApplication()

	err := core.InitSeeder(app.DB)
	if err != nil {
		log.Fatal(err)
	}

	router.InitRoutes(app)

	app.Start()
}
