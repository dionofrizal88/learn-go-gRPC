package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/dionofrizal88/go-grpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient){
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil{
		log.Fatalf("Error while creating stram: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Dio"},
		{FirstName: "Agus"},
		{FirstName: "Nofrizal"},
	}

	// creating chanel
	waitc := make(chan struct{})

	// goroutine sending request
	go func (){
		for _, req := range reqs{
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func (){
		for {
			res, err := stream.Recv()

			if err ==io.EOF{
				break
			}
			if err != nil{
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Recived: %v\n", res.Result)
		}
		close(waitc)
	}()

	<- waitc
}