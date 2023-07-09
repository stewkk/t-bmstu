package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/stewkk/t-bmstu/pkg/errors"

	"github.com/stewkk/t-bmstu/internal/adapters/html"
	"github.com/stewkk/t-bmstu/internal/adapters/rest"
	"github.com/stewkk/t-bmstu/internal/adapters/testsystem/timus"
	"github.com/stewkk/t-bmstu/internal/app"
)

func main() {
	restServer := rest.Server{
		App: app.NewApp(&timus.WebProblemRepo{}, &timus.TestingSystem{}),
	}
	htmlServer := html.Server{
		App: app.NewApp(&timus.WebProblemRepo{}, &timus.TestingSystem{}),
	}

	e := echo.New()
	e.Renderer = html.NewTemplate("web/templates/*.html")
	e.HTTPErrorHandler = errors.ErrorHandler

	e.Use(middleware.CORS())

	restGroup := e.Group("", rest.TagMiddleware)
	htmlGroup := e.Group("", html.TagMiddleware)

	rest.RegisterHandlers(restGroup, &restServer)
	html.RegisterHandlers(htmlGroup, &htmlServer)

	e.Logger.Fatal(e.Start(":8080"))
}
