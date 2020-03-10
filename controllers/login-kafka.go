package controllers

import (
	"github.com/dabory/svc-abango/locals"
	"encoding/json"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
)

type LoginReturn struct {
	locals.AnswerBase
	locals.Login
}

type LoginController struct {
	abango.Controller
}

func (c *LoginController) EditRow() error {

	var v locals.Login
	var r LoginReturn
	if err := json.Unmarshal(c.Ctx.Ask.Body, &v); err == nil {
		r.UserId = v.UserId
		r.Password = v.Password
		r.SvcStatus = "200"
		r.SvcMsg = "Everything is OK"
		c.Data["json"] = r

		c.AnswerJson()

	} else {
		return e.MyErr("salrqladksjfl-Unmarshal", err, false)
	}

	return nil
}
