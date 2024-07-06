package main

import (
	"example1/templates"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
This example shows a basic setup for an API that responds with HTML.

Uses:
	`echo` for a networking framework;
	`templ` for templating;
	`air` for live reloading.
*/

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {

		component := templates.Index("Example 1")
		return templRenderer(c, http.StatusOK, component)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func templRenderer(c echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	err := t.Render(c.Request().Context(), buf)
	if err != nil {
		return err
	}

	return c.HTMLBlob(statusCode, buf.Bytes())
}
