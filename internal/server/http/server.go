package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/drhelius/grpc-demo-proto/account"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func Serve(wg *sync.WaitGroup, grpc_port string, http_port string) {
	defer wg.Done()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := account.RegisterAccountServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", grpc_port), opts)
	if err != nil {
		return
	}

	log.Printf("Serving HTTP on localhost:%s ...", http_port)

	http.ListenAndServe(fmt.Sprintf(":%s", http_port), mux)
}
