package controllers

import (
	"encoding/json"

	lo "github.com/dabory/svc-abango/local"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
)

type LoginGrpcController struct {
	abango.Controller
}

func (c *LoginGrpcController) EditRow() (retBytes []byte, retSta []byte, err error) {
	var v lo.Login
	if err := json.Unmarshal(c.Ctx.Ask.Body, &v); err == nil {
		retBytes = c.Ctx.Ask.Body
		retSta = []byte("200")
	} else {
		return nil, nil, e.MyErr("wefarafsdfcvz-Unmarshal", err, false)
	}
	return retBytes, retSta, nil
}
