package main

import (
	"html/template"

	"github.com/labstack/echo/v4"

	"github.com/stewkk/t-bmstu/internal/api"
	"github.com/stewkk/t-bmstu/pkg/service"
	"github.com/stewkk/t-bmstu/pkg/views"
)

func main() {
	myApi := api.ServerInterfaceImpl{
		Service: service.TestingSystemService{},
	}
    e := echo.New()
	e.Renderer =  &views.Template{
		Templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}

    api.RegisterHandlers(e, &myApi)
	e.Logger.Fatal(e.Start(":8080"))
}
