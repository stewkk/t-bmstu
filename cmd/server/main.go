package main

import (
	"github.com/stewkk/t-bmstu/internal/api"

	"github.com/labstack/echo/v4"
)

func main() {
    var myApi api.ServerInterfaceImpl
    e := echo.New()
    api.RegisterHandlers(e, &myApi)
	e.Logger.Fatal(e.Start(":8080"))
}
