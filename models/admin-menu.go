package models

import (
	"dbrshop-svc/locals"
	"errors"

	e "github.com/dabory/abango/etc"
	"github.com/go-xorm/xorm"
)

type AdminMenu struct {
	Id           int64
	MenuNo       string `xorm:"varchar(10)"`
	Name         string `xorm:"varchar(60)"`
	Url          string `xorm:"varchar(200)"`
	Type         string `xorm:"varchar(10)"`
	Depth        int8
	locals.DbCom `xorm:"extends"`
}

func (m *AdminMenu) TableName() string { // 반드시 있어야 table name을 가져올 수 있다.
	return e.TableName(*m)
}

func (m *AdminMenu) AddRow(db *xorm.Engine) (int64, error) {

	if affected, err := db.Insert(m); affected == 1 && err == nil {
		return m.Id, nil
	} else {
		if err != nil {
			e.MyErr("ewrrwvczhu", err, false)
			return 0, err
		} else {
			e.ErrLog("sdfrtbbms-Record Not saved in "+m.TableName(), nil)
			return 0, errors.New("Record Not Added")
		}
	}

	return 0, nil
}

func (m *AdminMenu) EditRow(db *xorm.Engine) error {
	if affected, err := db.Id(m.Id).Update(m); affected == 1 && err == nil {
		return nil
	} else {
		if err != nil {
			e.MyErr("wefa34ferrqvcz", err, false)
			return err
		} else {
			e.ErrLog("fsdbvrers-Record Not updated in "+m.TableName(), nil)
			return errors.New("Record Not updated")
		}
	}

}

// func (t *UserDevice) DelRow() error {
// 	if affected, err := abango.XDB.Id(t.Id).Delete(t); affected == 1 && err == nil {
// 		return nil
// 	} else {
// 		if err != nil {
// 			e.MyErr("ewrvbhgfuryt", err, false)
// 			return err
// 		} else {
// 			e.ErrLog("sdfewdfafadf-Record Not deleted in "+t.TableName(), nil)
// 			return errors.New("Record Not deleted")
// 		}
// 	}

// }

type AdminMenuPrt struct {
	locals.AnswerBase
	locals.PageVars
	AdminMenus []AdminMenu
}

func AdminMenuPage(db *xorm.Engine, t *AdminMenuPrt) error {

	var cm AdminMenu
	if cnt, err := db.Where(t.PageVars.Query).Count(&cm); err == nil {
		t.PageVars.QueryCnt = cnt
	} else {
		return e.MyErr("rtret45687gfhk", err, false)
	}

	qry := *db.
		Where(t.PageVars.Query).
		Cols(t.PageVars.Fields).
		Limit(t.PageVars.Limit, t.PageVars.Offset)

	if t.PageVars.Asc != "" {
		qry = *qry.Asc(t.PageVars.Asc)
	}
	if t.PageVars.Desc != "" {
		qry = *qry.Desc(t.PageVars.Desc)
	}

	if err := qry.
		Find(&t.AdminMenus); err == nil {
		return nil
	} else {
		return e.MyErr("werqwrf234dgsg", err, false)
	}

}
