package clients

import (
	"log"

	"github.com/drhelius/grpc-demo-proto/user"
	"google.golang.org/grpc"
)

var UserService user.UserServiceClient

func init() {
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	UserService = user.NewUserServiceClient(conn)
}
