package controllers

import (
	"dbrshop-svc/locals"
	"encoding/json"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
)

type LoginRestController struct {
	abango.Controller
}

func (t *LoginRestController) EditRow() error {

	var v locals.Login
	if err := json.Unmarshal(t.Ctx.Ask.Body, &v); err == nil {
		// e.Tp(v)
		t.Ctx.Answer.Body = t.Ctx.Ask.Body
<<<<<<< HEAD
		// t.Ctx.Answer.Body. = []byte("200")
=======
>>>>>>> 552cb63d80ff6e09f916735103250c9e84c9a114
		// t.Answer()
	} else {
		return e.MyErr("salrqladksjfl-Unmarshal", err, false)
	}

	return nil
}
