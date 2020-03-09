package main

import (
	"dbrshop-svc/routers"

	"github.com/dabory/abango"
)

func init() {
	abango.RunServicePoint(routers.KafkaRouter, routers.GrpcConnect, routers.RestConnect) // Do not change !!
}

func main() {
}
