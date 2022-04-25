package main

import (
	ClientDepo "clientGRPC/resources/account"
	proto "clientGRPC/resources/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

const address = "localhost:9000"

var (
	i      int
	amount float32
)

func main() {
	x := 1

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := ClientDepo.NewDepositClient(conn)

	for x == 1 {
		fmt.Println("\n1. Deposit")
		fmt.Println("2. Check Balance")
		fmt.Println("0. Exit")
		fmt.Print(">>>>>")
		fmt.Scanln(&i)

		switch i {
		case 1:
			fmt.Println("\nEnter the amount you want to deposit.")
			fmt.Print(">>>>>")
			fmt.Scanln(&amount)

			res, err := client.Deposit(context.Background(), proto.DepositRequest{
				Amount: float32(amount),
			})

			if err != nil {
				log.Fatalf("Deposit Failed: %v", err)
			}
			if res == true {
				log.Println("Deposit Success")
				return
			}
		case 2:
			res, err := client.GetDeposit(context.Background())
			if err != nil {
				log.Fatalf("\nCheck Balance Failed: %v", err)
			}

			log.Printf("%v", res)
		case 0:
			log.Println("\n>>>>> THANK YOU <<<<<")
			x = 0
		default:
			log.Println("\nMenu Unavailable")
		}
	}
}
