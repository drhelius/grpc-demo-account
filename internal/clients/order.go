package clients

import (
	"log"

	"github.com/drhelius/grpc-demo-proto/order"
	"google.golang.org/grpc"
)

var OrderService order.OrderServiceClient

func init() {
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	OrderService = order.NewOrderServiceClient(conn)
}
