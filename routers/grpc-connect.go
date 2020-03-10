package routers

import (
	"net"

	// grp1 "dbrshop-svc/protos"
	grp1 "github.com/dabory/abango/protos"

	"github.com/dabory/abango"
	e "github.com/dabory/abango/etc"
	"google.golang.org/grpc"
)

func GrpcConnect() {

	lis, err := net.Listen(abango.XConfig["gRpcProtocol"], abango.XConfig["gRpcAddr"]+":"+abango.XConfig["gRpcPort"])
	if err != nil {
		e.MyErr("QERQDFGVXGJER-failed to listen in gRpc", err, true)
	}

	grpcServer := grpc.NewServer()
	grp1.RegisterGrp1Server(grpcServer, NewGrp1Server())
	grpcServer.Serve(lis)
}

type grp1Server struct {
}

func NewGrp1Server() *grp1Server {
	return &grp1Server{}
}
