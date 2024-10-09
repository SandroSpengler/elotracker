package main

import (
	"io"
	"net/http"

	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Count struct {
	Count int
}

func newTemplate() *Template {
	t := Template{}

	t.templates = template.Must(template.ParseGlob("src/**/*.html"))

	return &t
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	count := Count{
		Count: 0,
	}

	e.Renderer = newTemplate()

	e.Static("/assets", "assets")

	e.GET("/", func(c echo.Context) error {
		count.Count++
		return c.Render(http.StatusOK, "index", count)
	})

	e.Logger.Fatal(e.Start(":5555"))
}
