package echo_web

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func JsonStringHandler(c echo.Context) error {
	u := &User{
		Name:  "Hello",
		Email: "World",
	}
	return c.JSON(http.StatusOK, u)
}

func HtmlStringHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
}

func WriteCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "Hello"
	cookie.Value = "World"
	cookie.Expires = time.Now().Add(60 * time.Second)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}
