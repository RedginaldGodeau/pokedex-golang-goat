package router

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

type Router struct {
	*echo.Echo
	Port string
}

func NewRouter(port string) *Router {
	return &Router{
		Echo: echo.New(),
		Port: ":" + port,
	}
}

func (r *Router) Serve() error {
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("FRONTEND")},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
	}))

	r.Static("/public", "website/public")
	err := r.Start(r.Port)
	if err != nil {
		return err
	}

	return nil
}

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
