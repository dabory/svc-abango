package local

import (
	e "github.com/dabory/abango/etc"
)

func (t *PageVars) ChkPageVars(chk string) {

	if t.Query == "" {
		e.ChkLog(chk, "Query is empty")
	}

	if t.Fields == "" {
		e.ChkLog(chk, "Fields is empty")
	}

	if t.Asc == "" {
		e.ChkLog(chk, "Asc is empty")
	}

	if t.Desc == "" {
		e.ChkLog(chk, "Desc is empty")
	}

	if t.Limit == 0 {
		e.ChkLog(chk, "Limit is zero")
	}

	if t.Offset == 0 {
		e.ChkLog(chk, "Offset is zero")
	}

	e.ChkLog(chk+", PageVars value is", *t)

	return
}
