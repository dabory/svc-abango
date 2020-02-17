package controllers

import (
	"encoding/json"

	lo "github.com/dabory/svc-abango/local"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
)

type LoginController struct {
	abango.Controller
}

func (c *LoginController) EditRow() error {

	var v lo.Login
	if err := json.Unmarshal(c.Ctx.Ask.Body, &v); err == nil {
		// e.Tp(v)
		c.Ctx.Answer.Body = c.Ctx.Ask.Body
		c.Ctx.Answer.Status = []byte("200")
		c.KafkaAnswer()
		// return
	} else {
		return e.MyErr("salrqladksjfl-Unmarshal", err, false)
	}

	return nil
}
