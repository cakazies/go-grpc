package routes

import (
	"context"
	"log"
	"os"

	rpc "github.com/cakazies/go-grpc/client/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type (
	// GrpcClient struct for running this file
	GrpcClient struct{}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Run function for running this client
func (g *GrpcClient) Run() {
	pro := serviceStudent()
	empty := rpc.Empty{}
	res3, err := pro.Get(context.Background(), &empty)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res3)
}

func serviceStudent() rpc.StudentsClient {
	port := os.Getenv("APPS_PORT")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return rpc.NewStudentsClient(conn)
}
