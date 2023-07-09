package html

import "github.com/labstack/echo/v4"

func TagMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Set("tag", "view")
		return next(ctx)
	}
}
