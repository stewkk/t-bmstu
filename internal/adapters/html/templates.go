package html

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
    Templates *template.Template
}

func NewTemplate(path string) *Template {
	return &Template{
		Templates: template.Must(template.ParseGlob(path)),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
