package controllers

import (
	"dbrshop-svc/locals"
	"encoding/json"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
)

type NotFoundController struct {
	abango.Controller
}

func (c *NotFoundController) NotFound() error {

	var r locals.AnswerBase
	if err := json.Unmarshal(c.Ctx.Ask.Body, &r); err == nil {
		r.SvcSts = "404"
		r.SvcMsg = "AskName [" + c.Ctx.Ask.AskName + "] Not Found !"
		c.Data["json"] = r
		c.AnswerJson()
	} else {
		return e.MyErr("salew63rqladksjfl-Unmarshal", err, false)
	}

	return nil
}
