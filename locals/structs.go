// Author : Eric Kim
// Build Date : 23 Jul 2008  Last Update 02 Aug 2008
// End-Agent for Passcon Multi OS go binding with Windows, MacOS, iOS, and Android
// All rights are reserved.

package locals

// 1. DB /////////////////////////////////////////////////////////////////

type DbCom struct {
	StartTime  int64 `xorm:"created"`
	EndTime    int64
	UpdateTime int64 `xorm:"updated"`
	Ip         string
	Usr        int64
}

// 2. Ask  /////////////////////////////////////////////////////////////////

type PageVars struct {
	QueryCnt       int64
	Query          string
	Fields         string
	Asc            string
	Desc           string
	Limit          int
	Offset         int
	InputFieldJson string
}

type InputFieldJson struct {
	StartDate  string
	EndDate    string
	Order      int
	Select     int
	QueryInput string
}

type Login struct {
	UserId   string
	Password string
}

// 3. Answer /////////////////////////////////////////////////////////////////
type AnswerBase struct {
	SvcStatus string
	SvcMsg    string
}

type ActRowCom struct {
	AnswerBase
	ActIdBase
}

type ActPageCom struct {
	AnswerBase
	ActIdsBase
}

type ActIdBase struct {
	Id int64
}

type ActIdsBase struct {
	Ids []struct {
		Id int64
	}
}
