package routers

import (
	"dbrshop-svc/controllers"

	"github.com/dabory/abango"
)

func KafkaRouter(ask *abango.AbangoAsk) {

	askname := ask.AskName
	if askname == "login" {
		var t controllers.LoginController
		t.Init(*ask)
		t.EditRow()
	} else if askname == "admin-menu-actrow" {
		var t controllers.AdminMenuController
		t.Init(*ask)
		t.ActRow()
	} else if askname == "admin-menu-page" {
		var t controllers.AdminMenuController
		t.Init(*ask)
		t.GetPage()
	}

}
