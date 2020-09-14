package grpc

import (
	"log"
	"net"
	"sync"

	"github.com/drhelius/grpc-demo-account/internal/impl"
	"github.com/drhelius/grpc-demo-proto/account"
	"google.golang.org/grpc"
)

func Serve(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("[Account] GRPC failed to listen: %v", err)
	}

	s := grpc.NewServer()

	account.RegisterAccountServiceServer(s, &impl.Server{})

	log.Printf("[Account] Serving GRPC on localhost:%s ...", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("[Account] GRPC failed to serve: %v", err)
	}
}
