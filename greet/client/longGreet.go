package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/dionofrizal88/go-grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient){
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest {
		{FirstName: "Dio"},
		{FirstName: "Agus"},
		{FirstName: "Nofrizal"},
	} // array kiriman request

	stream, err := c.LongGreet(context.Background()) // define endpoin yang ingin digunakan

	if err!= nil{
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs{
		fmt.Printf("Sending req: %v\n", req)

		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil{
		log.Fatalf("Error while reacieve response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}