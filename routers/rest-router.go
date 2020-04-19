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

	// // To show RequestBody Contents. !!! Use Json Validator First
	// m := echo.Map{}
	// if err := c.Bind(&m); err != nil {
	// 	return err
	// }
	// return c.JSON(200, m)

	err := json.NewDecoder(c.Request().Body).Decode(&v)
	if err != nil {
		return c.String(300, "QERQDFGVXGJER-Failed in Rest Request Body: "+
			echo.NewHTTPError(http.StatusInternalServerError, err.Error).Error())
	}

	askname := v.AskName
	e.OkLog("REST API [" + askname + "] arrived from " + c.RealIP())

	if askname == "login" {
		var t controllers.LoginRestController
		t.Init(*v)
		t.EditRow()
		return c.String(http.StatusOK, string(t.Ctx.Answer.Body))
	} else {
		var t controllers.NotFoundController
		t.Init(*v)
		t.NotFound()
	}
	return nil
}
