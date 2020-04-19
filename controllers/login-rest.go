package controllers

import (
	"encoding/json"

	"github.com/dabory/svc-abango/locals"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
)

type LoginRestController struct {
	abango.Controller
}

func (t *LoginRestController) EditRow() error {

	var v locals.Login
	var r LoginReturn
	if err := json.Unmarshal(t.Ctx.Ask.Body, &v); err == nil {

		r.UserId = v.UserId
		r.Password = v.Password
		r.SvcStatus = "200"
		r.SvcMsg = "Everything is OK"

		ret, _ := json.MarshalIndent(r, "", "  ")
		t.Ctx.Answer.Body = ret // t.Ctx.Answer.Body. = []byte("200")
		// ret, _ := json.Marshal(r)

	} else {
		return e.MyErr("salrqladksjfl-Unmarshal", err, false)
	}

	return nil
}
