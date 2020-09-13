package clients

import (
	"log"

	"github.com/drhelius/grpc-demo-proto/order"
	"google.golang.org/grpc"
)

var OrderService order.OrderServiceClient

func init() {
	log.Printf("[Account] Dialing to 'order:5000' ...")

	conn, err := grpc.Dial("order:5000", grpc.WithInsecure(), grpc.WithBlock(), grpc.FailOnNonTempDialError(true))
	if err != nil {
		log.Fatalf("[Account] Error dialing to Order service: %v", err)
	}

	OrderService = order.NewOrderServiceClient(conn)
}
