package clients

import (
	"log"

	"github.com/drhelius/grpc-demo-proto/user"
	"google.golang.org/grpc"
)

var UserService user.UserServiceClient

func init() {
	conn, err := grpc.Dial("user:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("[Account] User service did not connect: %v", err)
	}

	UserService = user.NewUserServiceClient(conn)
}
