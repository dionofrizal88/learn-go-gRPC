package main

import (
	"context"
	"io"
	"log"

	pb "github.com/dionofrizal88/go-grpc/greet/proto"
)

func greetManyTimes(c pb.GreetServiceClient){
	log.Println("GreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Dio",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil{
		log.Fatalf("Error while GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF{
			break
		}

		if err !=nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}