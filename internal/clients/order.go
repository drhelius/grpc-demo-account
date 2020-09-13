package clients

import (
	"log"
	"time"

	"github.com/drhelius/grpc-demo-proto/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var OrderService order.OrderServiceClient

func init() {
	log.Printf("[Account] Dialing to 'order:5000' ...")

	keepAliveParams := keepalive.ClientParameters{
		Time:                5 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}

	conn, err := grpc.Dial("order:5000", grpc.WithInsecure(), grpc.WithBlock(), grpc.FailOnNonTempDialError(true), grpc.WithKeepaliveParams(keepAliveParams))
	if err != nil {
		log.Fatalf("[Account] Error dialing to Order service: %v", err)
	}

	OrderService = order.NewOrderServiceClient(conn)
}
