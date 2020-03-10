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
		if id, err := m.AddRow(c.Db); err == nil {
			r.SvcStatus = "200"
			r.Id = id
		} else {
			r.SvcStatus = "604"
			r.SvcMsg = e.MyErr("rfadsywtsf-Record not Acted !", err, true).Error()
		}

	} else {
		r.SvcStatus = "601"
		r.SvcMsg = e.MyErr("wer324fgsvz-Record not Acted !", err, true).Error()
	}

	c.Data["json"] = r
	c.AnswerJson()
}

// func (c *AdminMenuController) EditRow() error {

// 	var v locals.Login
// 	if err := json.Unmarshal(c.Ctx.Ask.Body, &v); err == nil {
// 		c.Ctx.Answer.Body = c.Ctx.Ask.Body
// 		c.Ctx.Answer.Status = []byte("200")
// 		c.KafkaAnswer()
// 	} else {
// 		return e.MyErr("salrqladksjfl-Unmarshal", err, false)
// 	}

// 	return nil
// }

// func (c *AdminMenuController) ActRow(askstr string) (string, string, error) {

// 	var retstr []byte
// 	var v models.AdminMenu
// 	var retid locals.ReturnId

// 	if err := json.Unmarshal([]byte(askstr), &v); err != nil {
// 		return "", "601", e.MyErr("wergfjdfgrqwe", err, false)
// 	}

// 	if v.Id == 0 {
// 		if _, err := v.AddRow(); err == nil {
// 			retstr = nil
// 		} else {
// 			sess.Rollback()
// 			errStr := "Record [" + strconv.Itoa(int(v.Id)) + "] of Device_log not Added, Rollingback all requests"
// 			e.ErrLog(errStr, err)
// 			return "606", errStr, nil
// 		}
// 	} else if v.Id > 0 {
// 		if err := v.EditRow(); err == nil {
// 			retstr = nil
// 		} else {
// 			sess.Rollback()
// 			errStr := "Record [" + strconv.Itoa(int(v.Id)) + "] of Device_log not update, Rollingback all requests"
// 			e.ErrLog(errStr, err)
// 			return "606", errStr, nil
// 		}
// 	} else if v.Id < 0 {
// 		v.Id = -v.Id
// 		if err := v.DelRow(); err == nil {
// 			retstr = nil
// 		} else {
// 			sess.Rollback()
// 			errStr := "Record [" + strconv.Itoa(int(v.Id)) + "] of Device_log not deleted, Rollingback all requests"
// 			e.ErrLog(errStr, err)
// 			return "606", errStr, nil
// 		}
// 	}

// 	if id, err := v.AddRow(); err == nil {
// 		addedid.Id = id
// 		retstr, _ = json.Marshal(addedid)
// 		return "201", string(retstr), nil
// 	} else {
// 		return "604", "Record not found", nil
// 	}
// }
// }
