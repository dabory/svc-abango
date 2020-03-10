package routers

import (
	"github.com/dabory/svc-abango/controllers"

	"github.com/dabory/abango"
)

func KafkaRouter(ask *abango.AbangoAsk) {

	askname := ask.AskName
	if askname == "login" {
		var t controllers.LoginController
		t.Init(*ask)
		t.EditRow()
<<<<<<< HEAD
=======
	} else if askname == "admin-menu-actrow" {
		var t controllers.AdminMenuController
		t.Init(*ask)
		t.ActRow()
	} else if askname == "admin-menu-page" {
		var t controllers.AdminMenuController
		t.Init(*ask)
		t.GetPage()
>>>>>>> 5dab4c60ac9d8d468ee574acb61549aae4bd6200
	}

}
