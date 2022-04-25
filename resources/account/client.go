package account

import (
	proto "clientGRPC/resources/proto"
	"context"
	"time"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
)

type ClientDepo struct {
	conn   grpc.ClientConn
	client proto.DepositServiceClient
}

func NewDepositClient(conn *grpc.ClientConn) *ClientDepo {
	return &ClientDepo{
		conn:   grpc.ClientConn{},
		client: proto.NewDepositServiceClient(conn),
	}
}

func (dc *ClientDepo) Deposit(ctx context.Context, in interface{}) (interface{}, error) {
	var request *proto.DepositRequest
	err := mapstructure.Decode(in, &request)

	if err != nil {
		return nil, err
	}

	ctxOut, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	data, err := dc.client.Deposit(ctxOut, request)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dc *ClientDepo) GetDeposit(ctx context.Context) (interface{}, error) {
	ctxOut, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	data, err := dc.client.GetDeposit(ctxOut, &proto.GetDepositRequest{})
	if err != nil {
		return nil, err
	}

	return data, nil
}
