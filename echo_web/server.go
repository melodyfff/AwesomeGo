package echo_web

import "github.com/labstack/echo"

func Start() {
	e := echo.New()

	Bind(*e)

	e.Logger.Fatal(e.Start(":1323"))
}
