package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/stewkk/t-bmstu/internal/api"
	"github.com/stewkk/t-bmstu/internal/errors"
	"github.com/stewkk/t-bmstu/internal/views"
	"github.com/stewkk/t-bmstu/pkg/service"
)

func main() {
	apiImpl := api.ServerInterfaceImpl{
		Service: service.TestingSystemService{},
	}
	viewsImpl := views.ServerInterfaceImpl{
		Service: service.TestingSystemService{},
	}

    e := echo.New()
	e.Renderer =  &views.Template{
		Templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}
	e.HTTPErrorHandler = errors.ErrorHandler

	e.Use(middleware.CORS())

	apiGroup := e.Group("", api.TagApiMiddleware)
	viewGroup := e.Group("", views.TagViewMiddleWare)

    api.RegisterHandlers(apiGroup, &apiImpl)
	views.RegisterHandlers(viewGroup, &viewsImpl)

	e.Logger.Fatal(e.Start(":8080"))
}
