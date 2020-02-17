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
	} else {
	}

}
