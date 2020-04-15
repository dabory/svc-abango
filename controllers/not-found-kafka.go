package controllers

import (
	"github.com/dabory/abango"
	"github.com/dabory/svc-abango/locals"
)

type NotFoundController struct {
	abango.Controller
}

func (c *NotFoundController) NotFound() error {

	var r locals.AnswerBase
	// if err := json.Unmarshal(c.Ctx.Ask.Body, &r); err == nil {
	// 	r.SvcStatus = "404"
	// 	r.SvcMsg = "AskName [" + c.Ctx.Ask.AskName + "] Not Found !"
	// 	c.Data["json"] = r
	// 	c.AnswerJson()
	// } else {
	// 	return e.MyErr("salew63rqladksjfl-Unmarshal", err, false)
	// }

	r.SvcStatus = "202"
	r.SvcMsg = "This is Answer"
	c.Data["json"] = r
	c.AnswerJson()

	return nil
}
