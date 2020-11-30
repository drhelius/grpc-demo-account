package clients

import (
	"log"

	"github.com/drhelius/grpc-demo-proto/user"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

var UserService user.UserServiceClient

func init() {
	log.Printf("[Account] Dialing to 'user:5000' ...")

	conn, err := grpc.Dial("user:5000", grpc.WithInsecure(), grpc.WithBlock(), grpc.FailOnNonTempDialError(true), grpc.WithStreamInterceptor(
		grpc_opentracing.StreamClientInterceptor(
			grpc_opentracing.WithTracer(opentracing.GlobalTracer()))))
	if err != nil {
		log.Fatalf("[Account] Error dialing to User service: %v", err)
	}

	UserService = user.NewUserServiceClient(conn)
}
