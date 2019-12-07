package main

import (
	"github.com/cakazies/go-grpc/client/routes"
)

func main() {
	client := routes.GrpcClient{}
	client.Run()
}
