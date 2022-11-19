package main

import (
	"context"
	"log"

	pb "github.com/dionofrizal88/go-grpc/calculator/proto"
)

func twoSum(c pb.CalculatorServiceClient){
	log.Println("twoSum was invoked")

	res, err := c.Sum(context.Background(), &pb.CalculatorRequest{
		OperanOne: 3,
		OperanTwo: 10,
	})

	if err != nil{
		log.Fatalf("Coul not sum: %v\n", err)
	}

	log.Printf("The result is: %d\n", res.Result)
}
