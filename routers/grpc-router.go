package routers

import (
	"context"
	"encoding/json"

	"github.com/dabory/svc-abango/controllers"

	grp1 "github.com/dabory/abango/protos"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
	"google.golang.org/grpc/peer"
)

func (s *grp1Server) StdRpc(c context.Context, ask *grp1.StdAsk) (*grp1.StdRet, error) {

	retbytes := []byte("")
	retsta := []byte("")
	var err error

	p, _ := peer.FromContext(c)
	ip_no_serial := p.Addr.String()
	// arr_ip := strings.Split(ip_no, ":")

	var v *abango.AbangoAsk
	if err := json.Unmarshal(ask.AskMsg, &v); err != nil {
		return nil, e.MyErr("RQCVSWDZVCXER-Unmarshall-ask.AskMsg", err, false)
	}

	askname := v.AskName
	e.OkLog("gRpc-API [" + askname + "] arrived from " + ip_no_serial)

	if askname == "login" {
		var t controllers.LoginGrpcController
		t.Init(*v)
		retbytes, retsta, err = t.EditRow()
	} else {
		var t controllers.NotFoundController
		t.Init(*v)
		t.NotFound()
	}

	ret := grp1.StdRet{
		RetMsg: retbytes,
		RetSta: retsta,
	}

	return &ret, err

}
