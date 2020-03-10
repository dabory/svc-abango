package controllers

import (
	"dbrshop-svc/locals"
	"dbrshop-svc/models"
	"encoding/json"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
)

type AdminMenuController struct {
	abango.Controller
}

func (c *AdminMenuController) GetPage() {

	var r models.AdminMenuPrt
	if err := json.Unmarshal(c.Ctx.Ask.Body, &r.PageVars); err == nil {
		if err := models.AdminMenuPage(c.Db, &r); err == nil {
			r.SvcStatus = "200"
		} else {
			r.SvcStatus = "601"
			r.SvcMsg = e.MyErr("ewrt45hgfh-Record not Acted !", err, true).Error()
		}
	} else {

	}
	c.Data["json"] = r
	c.AnswerJson()
}

func (c *AdminMenuController) ActRow() {

	var r locals.ActRowCom
	var m models.AdminMenu

	if err := json.Unmarshal(c.Ctx.Ask.Body, &m); err == nil {
		chkid := m.Id
		if chkid == 0 {
			if id, err := m.AddRow(c.Db); err == nil {
				r.SvcStatus = "200"
				r.Id = id
			} else {
				r.SvcStatus = "603"
				r.SvcMsg = err.Error()
			}
		} else if chkid > 0 {
			if err := m.EditRow(c.Db); err == nil {
				r.SvcStatus = "200"
				r.Id = chkid
			} else {
				r.SvcStatus = "605"
				r.Id = chkid
				r.SvcMsg = err.Error()
			}
		} else if chkid < 0 {
			absid := -chkid
			m.Id = absid
			if err := m.DelRow(c.Db); err == nil {
				r.SvcStatus = "200"
				r.Id = absid
			} else {
				r.SvcStatus = "606"
				r.SvcMsg = err.Error()
			}
		}

	} else {
		r.SvcStatus = "601"
		r.SvcMsg = e.MyErr("wer324fgsvz-Record not Acted !", err, true).Error()
	}

	c.Data["json"] = r
	c.AnswerJson()
}
