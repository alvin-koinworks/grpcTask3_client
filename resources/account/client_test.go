package account_test

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"

	"clientGRPC/resources/account"
	proto "clientGRPC/resources/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type mockDepositServiceServer struct {
	proto.UnimplementedDepositServiceServer
}

type Deposit struct {
	Amount float32
}

var Dep Deposit

func dial() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	proto.RegisterDepositServiceServer(server, &mockDepositServiceServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func (d *mockDepositServiceServer) Deposit(ctx context.Context, in *proto.DepositRequest) (*proto.DepositResponse, error) {
	// log.Printf("mock deposit-1: %v", Dep)
	if in.GetAmount() < 10000 {
		return nil, status.Errorf(codes.InvalidArgument, "cannot deposit %v", in.GetAmount())
	}

	Dep.Amount = Dep.Amount + in.GetAmount()

	// log.Printf("mock deposit: %v", Dep)
	return &proto.DepositResponse{Ok: true}, nil

}

func (d *mockDepositServiceServer) GetDeposit(ctx context.Context, in *proto.GetDepositRequest) (*proto.GetDepositResponse, error) {
	getResponse := proto.GetDepositResponse{
		TotalAmount: Dep.Amount,
	}

	return &getResponse, nil
}

func TestDepositServiceClient_Deposit(t *testing.T) {
	test := []struct {
		name   string
		amount float32
		res    bool
		err    string
	}{
		{
			"Invalid Request with Invalid Amount",
			5000,
			false,
			"deposit value must be at least 10000",
		},
		// {
		// 	"Valid Request with Valid Amount",
		// 	10000,
		// 	true,
		// 	"",
		// },
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dial()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			var value bool
			request := &proto.DepositRequest{Amount: tc.amount}
			ctx := context.Background()

			client := account.NewDepositClient(conn)
			response, err := client.Deposit(ctx, request)

			// log.Printf("amount: %v", tc.amount)

			// log.Printf("request: %v", request)

			// log.Printf("response: %v", response)

			// log.Printf("value: %v", value)

			if response != tc.res {
				log.Print(tc)
				log.Print(response)
				t.Error("error: expected", tc.res, "received", response)
			}

			if err != nil && errors.Is(err, err) {
				t.Error("error: expected", tc.err, "received", err)
			}

			if response != nil {
				val := response.(*proto.DepositResponse)
				value = val.Ok
			}

			if value != tc.res {
				t.Error("error: expected", tc.res, "received:", value)
			}

			if err != nil {
				if er, _ := status.FromError(err); er.Message() != tc.err {
					t.Error("Error message: ", tc.err, "received", er.Message())
				}
			}
		})
	}
}

func TestDepositServiceClient_GetDeposit(t *testing.T) {
	test := struct {
		name string
		res  *proto.GetDepositResponse
		err  string
	}{
		"Valid Test GetDeposit",
		&proto.GetDepositResponse{TotalAmount: 10000},
		"",
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dial()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	t.Run(test.name, func(t *testing.T) {
		client := account.NewDepositClient(conn)
		getResponse, err := client.GetDeposit(context.Background())
		val := getResponse.(*proto.GetDepositResponse)
		if val.GetTotalAmount() != test.res.GetTotalAmount() {
			t.Error("error : expected", test.res, "received:", val)
		}
		if err != nil {
			if er, _ := status.FromError(err); er.Message() != test.err {
				t.Error("Error message: ", test.err, "received", er.Message())
			}
		}
	})
}
