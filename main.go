package main

import (
	"fmt"

	"github.com/dabory/svc-abango/routers"

	"github.com/dabory/abango"
)

func init() {
	fmt.Print("!! Notice: svc-abango should be run in $GOSRC/github.com/dabory/svc-abango/")
	abango.RunServicePoint(routers.KafkaRouter, routers.GrpcConnect, routers.RestConnect) // Do not change !!
}

func main() {
}
