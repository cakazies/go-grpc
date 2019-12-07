package main

import (
	"github.com/cakazies/go-grpc/server/routes"
)

func main() {
	grpc := routes.GrpcRoute{}
	grpc.Run()
}
