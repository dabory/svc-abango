package controllers

import (
	"encoding/json"

	lo "github.com/dabory/svc-abango/local"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
)

type LoginRestController struct {
	abango.Controller
}

func (t *LoginRestController) EditRow() error {

	var v lo.Login
	if err := json.Unmarshal(t.Ctx.Ask.Body, &v); err == nil {
		// e.Tp(v)
		t.Ctx.Answer.Body = t.Ctx.Ask.Body
		// t.Ctx.Answer.Body. = []byte("200")
		// t.Answer()
	} else {
		return e.MyErr("salrqladksjfl-Unmarshal", err, false)
	}

	return nil
}
