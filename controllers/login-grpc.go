package controllers

import (
	"encoding/json"

	"github.com/dabory/svc-abango/locals"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
)

type LoginGrpcController struct {
	abango.Controller
}

func (c *LoginGrpcController) EditRow() (retBytes []byte, retSta []byte, err error) {

	var v locals.Login
	var r LoginReturn
	if err := json.Unmarshal(c.Ctx.Ask.Body, &v); err == nil {

		r.UserId = v.UserId
		r.Password = v.Password
		r.SvcStatus = "200"
		r.SvcMsg = "Everything is OK"

		ret, _ := json.MarshalIndent(r, "", "  ")
		// ret, _ := json.Marshal(r)

		retBytes = ret
		retSta = []byte("200")

	} else {
		return nil, nil, e.MyErr("wefarafsdfcvz-Unmarshal", err, false)
	}
	return retBytes, retSta, nil
}
