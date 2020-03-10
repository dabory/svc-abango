package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dabory/svc-abango/controllers"

	"github.com/dabory/abango"
	"github.com/labstack/echo"

	e "github.com/dabory/abango/etc"
)

func RestRouter(c echo.Context) error {

	var v *abango.AbangoAsk

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&v)
	if err != nil {
		return e.MyErr("QERQDFGVXGJER-Failed in Rest Request Body",
			echo.NewHTTPError(http.StatusInternalServerError, err.Error), true)
	}

	askname := v.AskName
	e.OkLog("REST API [" + askname + "] arrived from " + c.RealIP())

	if askname == "login" {
		var t controllers.LoginRestController
		t.Init(*v)
		t.EditRow()
		return c.String(http.StatusOK, string(t.Ctx.Answer.Body))
	} else {
		return nil
	}
	return nil
}
