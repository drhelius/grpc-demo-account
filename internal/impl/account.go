package impl

import (
	"context"
	"log"
	"time"

	"github.com/drhelius/grpc-demo-account/internal/clients"
	"github.com/drhelius/grpc-demo-proto/account"
	"github.com/drhelius/grpc-demo-proto/order"
	"github.com/drhelius/grpc-demo-proto/user"
)

type Server struct {
	account.UnimplementedAccountServiceServer
}

func (s *Server) Create(ctx context.Context, in *account.CreateAccountReq) (*account.CreateAccountResp, error) {

	log.Printf("Received: %s", in.GetAccount())

	return &account.CreateAccountResp{Id: "testid"}, nil
}

func (s *Server) Read(ctx context.Context, in *account.ReadAccountReq) (*account.ReadAccountResp, error) {

	log.Printf("Received: %v", in.GetId())

	u := getUser(in.GetId())
	o := getOrder(in.GetId())
	orders := []*order.Order{o}

	return &account.ReadAccountResp{Account: &account.Account{Id: "demoid", User: u, Orders: orders}}, nil
}

func getUser(id string) *user.User {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	u, err := clients.UserService.Read(ctx, &user.ReadUserReq{Id: id})

	if err != nil {
		log.Fatalf("Could not invoke User service: %v", err)
	}

	log.Printf("User service invocation: %v", u.GetUser())

	return u.GetUser()
}

func getOrder(id string) *order.Order {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	o, err := clients.OrderService.Read(ctx, &order.ReadOrderReq{Id: id})

	if err != nil {
		log.Fatalf("Could not invoke Order service: %v", err)
	}

	log.Printf("Order service invocation: %v", o.GetOrder())

	return o.GetOrder()
}
