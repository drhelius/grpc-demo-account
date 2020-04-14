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
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	account.RegisterAccountServiceServer(s, &impl.Server{})

	log.Printf("Serving GRPC on localhost:%s ...", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
