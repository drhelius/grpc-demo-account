package clients

import (
	"log"
	"time"

	"github.com/drhelius/grpc-demo-proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var UserService user.UserServiceClient

func init() {
	log.Printf("[Account] Dialing to 'user:5000' ...")

	keepAliveParams := keepalive.ClientParameters{
		Time:                5 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}

	conn, err := grpc.Dial("user:5000", grpc.WithInsecure(), grpc.WithBlock(), grpc.FailOnNonTempDialError(true), grpc.WithKeepaliveParams(keepAliveParams))
	if err != nil {
		log.Fatalf("[Account] Error dialing to User service: %v", err)
	}

	UserService = user.NewUserServiceClient(conn)
}
