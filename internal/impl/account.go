package impl

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/drhelius/grpc-demo-account/internal/clients"
	"github.com/drhelius/grpc-demo-proto/account"
	"github.com/drhelius/grpc-demo-proto/order"
	"github.com/drhelius/grpc-demo-proto/user"
)

type Server struct {
	account.UnimplementedAccountServiceServer
}

func (s *Server) Create(ctx context.Context, in *account.CreateAccountReq) (*account.CreateAccountResp, error) {

	log.Printf("[Account] Create Req: %v", in.GetAccount())

	r := &account.CreateAccountResp{Id: strconv.Itoa(randomdata.Number(1000000))}

	log.Printf("[Account] Create Res: %v", r.GetId())

	return r, nil
}

func (s *Server) Read(ctx context.Context, in *account.ReadAccountReq) (*account.ReadAccountResp, error) {

	log.Printf("[Account] Read Req: %v", in.GetId())

	u := getUser(ctx, strconv.Itoa(randomdata.Number(1000000)))
	o1 := getOrder(ctx, strconv.Itoa(randomdata.Number(1000000)))
	o2 := getOrder(ctx, strconv.Itoa(randomdata.Number(1000000)))
	o3 := getOrder(ctx, strconv.Itoa(randomdata.Number(1000000)))

	orders := []*order.Order{o1, o2, o3}

	r := &account.ReadAccountResp{Account: &account.Account{Id: in.GetId(), User: u, Orders: orders}}

	log.Printf("[Account] Read Res: %v", r.GetAccount())

	return r, nil
}

func getUser(ctx context.Context, id string) *user.User {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	log.Printf("[Account] Invoking User service: %s", id)

	u, err := clients.UserService.Read(ctxTimeout, &user.ReadUserReq{Id: id})

	if err != nil {
		log.Printf("[Account] ERROR - Could not invoke User service: %v", err)
		return &user.User{}
	}

	log.Printf("[Account] User service invocation: %v", u.GetUser())
	return u.GetUser()
}

func getOrder(ctx context.Context, id string) *order.Order {

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	log.Printf("[Account] Invoking Order service: %s", id)

	o, err := clients.OrderService.Read(ctxTimeout, &order.ReadOrderReq{Id: id})

	if err != nil {
		log.Printf("[Account] ERROR - Could not invoke Order service: %v", err)
		return &order.Order{}
	}

	log.Printf("[Account] Order service invocation: %v", o.GetOrder())
	return o.GetOrder()
}
