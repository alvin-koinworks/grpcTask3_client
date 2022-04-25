package account_test

import (
	"context"
	"fmt"
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
		name    string
		amount  float32
		res     *proto.DepositResponse
		errCode codes.Code
		err     string
	}{
		{
			"Invalid Request with Invalid Amount",
			5000,
			nil,
			codes.InvalidArgument,
			fmt.Sprintf("cannot deposit %v", 5000),
		},
		{
			"Valid Request with Valid Amount",
			10000,
			&proto.DepositResponse{Ok: true},
			codes.OK,
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dial()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewDepositServiceClient(conn)

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			// var value bool
			request := &proto.DepositRequest{Amount: tc.amount}
			ctx := context.Background()

			response, err := client.Deposit(ctx, request)

			if response != nil {
				if response.GetOk() != tc.res.GetOk() {
					t.Error("response: expected", tc.res.GetOk(), "received", response.GetOk())
				}
			}

			if err != nil {
				if er, ok := status.FromError(err); ok {

					log.Printf("tc.err: %v", tc.err)

					log.Printf("er.Message: %v", er.Message())

					if er.Code() != tc.errCode {
						t.Error("error code: expected", codes.InvalidArgument, "received", er.Code())
					}
					if er.Message() != tc.err {
						t.Error("error message: expected", tc.err, "received", er.Message())
					}
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
