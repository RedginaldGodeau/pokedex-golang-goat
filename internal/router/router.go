package router

import (
	_ "backend/docs"
	"backend/internal/controller"
	"backend/pkg/application"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func InitRoutes(app *application.Application) {

	ViewsController := controller.NewViewsController()
	PokemonController := controller.NewPokemonController(app)

	app.Router.GET("/", ViewsController.Index)
	app.Router.GET("/:id/", ViewsController.PokemonByID)

	api := app.Router.Group("/api")
	{
		api.GET("/ping/", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})
		api.GET("/swagger/*", echoSwagger.WrapHandler)

		connexion := api.Group("/pokemon")
		{
			connexion.GET("/:id/", PokemonController.GetByID)
			connexion.GET("/", PokemonController.GetAll)
			connexion.GET("/:pokedex/pokedex/", PokemonController.GetByPokedex)
			connexion.GET("/:name/name/", PokemonController.GetByName)
			connexion.GET("/:search/search/", PokemonController.GetBySearch)
		}
	}

}
