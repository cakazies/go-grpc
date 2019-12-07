package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	rpc "github.com/cakazies/go-grpc/server/proto"
	"github.com/joho/godotenv"

	"google.golang.org/grpc"
)

type (
	// GrpcRoute struct for running this file
	GrpcRoute struct{}
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file : ", err)
	}
}

// Run function for running this server
func (g *GrpcRoute) Run() {
	srv := grpc.NewServer()
	var garageSrv GrpcRoute
	rpc.RegisterStudentsServer(srv, garageSrv)
	port := os.Getenv("APPS_PORT")
	log.Println("Starting RPC server at", port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", port, err)
	}
	log.Fatal(srv.Serve(l))
}

// Get function for get all data
func (GrpcRoute) Get(ctx context.Context, req *rpc.Empty) (*rpc.Reponse, error) {
	var resp rpc.Reponse
	var student rpc.Student

	student.Id = 10
	student.Jk = "Women"
	student.Name = "Meliana Lily"
	student.Nilai = 90
	student.Nis = "012019018"
	fmt.Println("sapi")
	result, err := json.Marshal(&student)
	if err != nil {
		return &resp, err
	}

	resp.Data = string(result)
	return &resp, nil
}
