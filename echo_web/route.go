package echo_web

import (
	"github.com/labstack/echo"
	"net/http"
)

func Bind(e echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/json", JsonStringHandler)

	e.GET("/html", HtmlStringHandler)

	e.GET("/cookie", WriteCookie)
}
