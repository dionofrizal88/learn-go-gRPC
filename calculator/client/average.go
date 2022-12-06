package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/dionofrizal88/go-grpc/calculator/proto"
)

func average(c pb.CalculatorServiceClient){
	reqs := []*pb.FloatRequest{
		{Number: 2},
		{Number: 3},
		{Number: 2},
	}

	stream, err := c.Average(context.Background())

	if err != nil{
		log.Fatalf("Error while calling average: %v\n", err)
	}

	for _, req := range reqs{
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("Error while receive response from average: %v\n", err)
	}

	log.Printf("Average: %.2f\n", res.Res)
}

// 	for _, req := range reqs{
// 		fmt.Printf("Sending req: %v\n", req)

// 		stream.Send(req)
// 		time.Sleep(1 * time.Second)
// 	}

// 	res, err := stream.CloseAndRecv()

// 	if err != nil{
// 		log.Fatalf("Error while reacieve response from LongGreet: %v\n", err)
// 	}

// 	log.Printf("LongGreet: %s\n", res.Result)
// }