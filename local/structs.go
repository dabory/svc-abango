// Author : Eric Kim
// Build Date : 23 Jul 2008  Last Update 02 Aug 2008
// End-Agent for Passcon Multi OS go binding with Windows, MacOS, iOS, and Android
// All rights are reserved.

package local

// 1. Receivers /////////////////////////////////////////////////////////////////

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
	MacAddr  string // 이하는 app-login에서 사용함.
	// RDeviceYes bool
	// CDeviceYes bool
}

type AddedId struct {
	Id int64
}

type PageIds struct {
	Ids []struct {
		Id int64
	}
}

// 2. Returners /////////////////////////////////////////////////////////////////

// 3. RecOne Returners /////////////////////////////////////////////////////////////////
type UserSiteRecOne struct {
	Id         int    `json:"Id"`
	SiteUserNo string `json:"SiteUserNo"`
	Passwd     string `json:"Passwd"`
}

type UserVaultRecOne struct {
	Id       int    `json:"Id"`
	Contents string `json:"Contents"`
}

type UserDeviceRecOne struct {
	Id  int    `json:"Id"`
	Upb string `json:"Upb"`
	Dak string `json:"Dak"`
}

type UserDeviceRecHolder struct {
	Id   int    `json:"Id"`
	Upb  string `json:"Upb"`
	EDak string `json:"EDak"`
}

type UserDeviceRecBackuper struct {
	Id   int    `json:"Id"`
	EUpd string `json:"EUpd"`
}

type UserDeviceEUpdOne struct {
	Id   int    `json:"Id"`
	EUpd string `json:"EUpd"`
}
