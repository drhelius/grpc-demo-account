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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	account.UnimplementedAccountServiceServer
}

func (s *Server) Create(ctx context.Context, in *account.CreateAccountReq) (*account.CreateAccountResp, error) {

	log.Printf("[Account] Create Req: %v", in.GetAccount())

	r := &account.CreateAccountResp{Id: strconv.Itoa(randomdata.Number(1000000))}

	err := failedContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("[Account] Create Res: %v", r.GetId())

	return r, nil
}

func (s *Server) Read(ctx context.Context, in *account.ReadAccountReq) (*account.ReadAccountResp, error) {

	log.Printf("[Account] Read Req: %v", in.GetId())

	err := failedContext(ctx)
	if err != nil {
		return nil, err
	}

	u := getUser(ctx, strconv.Itoa(randomdata.Number(1000000)))

	var orders [3]*order.Order

	for i := 0; i < 3; i++ {
		err = failedContext(ctx)
		if err != nil {
			return nil, err
		}

		orders[i] = getOrder(ctx, strconv.Itoa(randomdata.Number(1000000)))
	}

	r := &account.ReadAccountResp{Account: &account.Account{Id: in.GetId(), User: u, Orders: orders[:]}}

	log.Printf("[Account] Read Res: %v", r.GetAccount())

	return r, nil
}

func getUser(ctx context.Context, id string) *user.User {

	log.Printf("[Account] Invoking User service: %s", id)

	headersIn, _ := metadata.FromIncomingContext(ctx)

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctxTimeout, headersIn)

	u, err := clients.UserService.Read(ctx, &user.ReadUserReq{Id: id})

	if err != nil {
		log.Printf("[Account] ERROR - Could not invoke User service: %v", err)
		return &user.User{}
	}

	log.Printf("[Account] User service invocation: %v", u.GetUser())
	return u.GetUser()
}

func getOrder(ctx context.Context, id string) *order.Order {

	log.Printf("[Account] Invoking Order service: %s", id)

	headersIn, _ := metadata.FromIncomingContext(ctx)

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	ctx = metadata.NewOutgoingContext(ctxTimeout, headersIn)

	o, err := clients.OrderService.Read(ctx, &order.ReadOrderReq{Id: id})

	if err != nil {
		log.Printf("[Account] ERROR - Could not invoke Order service: %v", err)
		return &order.Order{}
	}

	log.Printf("[Account] Order service invocation: %v", o.GetOrder())
	return o.GetOrder()
}

func failedContext(ctx context.Context) error {
	if ctx.Err() == context.Canceled {
		log.Printf("[Account] context canceled, stoping server side operation")
		return status.Error(codes.Canceled, "context canceled, stoping server side operation")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("[Account] dealine has exceeded, stoping server side operation")
		return status.Error(codes.DeadlineExceeded, "dealine has exceeded, stoping server side operation")
	}

	return nil
}
