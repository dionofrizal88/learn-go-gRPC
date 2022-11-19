package main

import (
	"context"
	"io"
	"log"

	pb "github.com/dionofrizal88/go-grpc/calculator/proto"
)

func sumFactor(c pb.CalculatorServiceClient) {
	log.Println("SumFactor was invoked")

	req := &pb.CalculatorRequest{
		OperanOne: 120,
		OperanTwo: 2,
	}

	stream, err := c.SumFactor(context.Background(), req)

	if err != nil{
		log.Fatalf("Error while SumFactor: WW%v\n", err)
	}

	for{
		msg, err := stream.Recv()

		if err == io.EOF{
			break
		}

		if err != nil{
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("The Primes: %s\n", msg.Result)
	}
}