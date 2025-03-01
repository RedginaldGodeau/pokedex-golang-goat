package controller

import (
	"backend/internal/context/pokemon"
	"backend/pkg/application"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type PokemonController struct {
	service pokemon.ServiceInterface
}

func NewPokemonController(app *application.Application) *PokemonController {
	return &PokemonController{
		service: pokemon.NewServiceFromApp(app),
	}
}

// @Summary Pokemon
// @Description Get Pokemon by ID
// @Tags pokemon
// @Accept  json
// @Produce  json
// @Param id path string true "Pokemon ID"
// @Success 200 {object} pokemon.PokemonResponse
// @Router /api/pokemon/{id}/ [get]
func (c *PokemonController) GetByID(ctx echo.Context) error {
	rawPokemonID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dto, err := c.service.GetByID(rawPokemonID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dto)
}

// @Summary Pokemon
// @Description Get Pokemon by Pokedex ID
// @Tags pokemon
// @Accept  json
// @Produce  json
// @Param pokedex path string true "Pokemon pokedex id"
// @Success 200 {array} pokemon.PokemonResponse
// @Router /api/pokemon/{pokedex}/pokedex/ [get]
func (c *PokemonController) GetByPokedex(ctx echo.Context) error {
	rawPokedexID, err := strconv.Atoi(ctx.Param("pokedex"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dto, err := c.service.GetByPokedex(rawPokedexID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dto)
}

// @Summary Pokemon
// @Description Get Pokemon by name
// @Tags pokemon
// @Accept  json
// @Produce  json
// @Param name path string true "Pokemon name"
// @Success 200 {object} pokemon.PokemonResponse
// @Router /api/pokemon/{name}/name/ [get]
func (c *PokemonController) GetByName(ctx echo.Context) error {
	rawName := ctx.Param("pokedex")
	if rawName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Pokedex name")
	}

	dto, err := c.service.GetByName(rawName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dto)
}

// @Summary Pokemon
// @Description Get all Pokemons
// @Tags pokemon
// @Accept  json
// @Produce  json
// @Success 200 {array} pokemon.PokemonResponse
// @Router /api/pokemon/ [get]
func (c *PokemonController) GetAll(ctx echo.Context) error {
	dto, err := c.service.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dto)
}

// @Summary Pokemon
// @Description Get Pokemon by name
// @Tags pokemon
// @Accept  json
// @Produce  json
// @Param search path string true "Pokemon name search"
// @Success 200 {array} pokemon.PokemonResponse
// @Router /api/pokemon/{search}/search/ [get]
func (c *PokemonController) GetBySearch(ctx echo.Context) error {
	rawSearch := ctx.Param("search")
	dto, err := c.service.GetBySearch(rawSearch)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, dto)
}
