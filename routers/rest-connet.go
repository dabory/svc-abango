package routers

import (
	"github.com/dabory/abango"
	"github.com/labstack/echo"
)

func RestConnect(ask *abango.AbangoAsk) {
	e := echo.New()
	e.POST("/", RestRouter) // RestUri is fixed

	// e.POST(abango.XConfig["RestUri"], func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World!")
	// })
	e.Logger.Fatal(e.Start(abango.XConfig["RestConnect"]))

	// e.Logger.Fatal(e.Start(":18080")) // localhost:1323

}
