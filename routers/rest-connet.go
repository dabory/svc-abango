package routers

import (
	"github.com/dabory/abango"
	"github.com/labstack/echo"
)

func RestConnect(ask *abango.AbangoAsk) {
	c := echo.New()
	c.POST(abango.XConfig["RestUri"], RestRouter) // RestUri is fixed

	c.Logger.Fatal(c.Start(abango.XConfig["RestConnect"]))
}
