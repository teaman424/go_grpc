package main

import (
	"context"
	"fmt"
	calculatorPB "grpcSample/proto/calculator"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := calculatorPB.NewCalculatorServiceClient(conn)

	doUnary(client)
}

func doUnary(client calculatorPB.CalculatorServiceClient) {
	fmt.Println("Staring to do a Unary RPC")

	req := &calculatorPB.CalculatorRequest{
		A: 23003,
		B: 394445,
	}

	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Sum failed : %v", err)
	}

	fmt.Println("Response from Server : ", res.Result)
}
