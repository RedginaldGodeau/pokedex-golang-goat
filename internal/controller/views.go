package controller

import (
	"backend/pkg/router"
	"backend/website/views"
	"backend/website/views/components/layouts"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ViewsController struct{}

func NewViewsController() *ViewsController {
	return &ViewsController{}
}

func (c *ViewsController) Index(ctx echo.Context) error {
	return router.Render(ctx, layouts.Layout(layouts.LayoutProps{Children: views.HomePage()}))
}

func (c *ViewsController) PokemonByID(ctx echo.Context) error {
	rawPokemonID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return router.Render(ctx, layouts.Layout(layouts.LayoutProps{Children: views.Pokemon(rawPokemonID)}))
}
