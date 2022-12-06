package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/dionofrizal88/go-grpc/calculator/proto"
)

func maxValue(c pb.CalculatorServiceClient){
	log.Println("maxValue was invoked")

	stream, err := c.Max(context.Background())

	if err != nil{
		log.Fatalf("Error while creating stream: %v\n", err)
	}
	
	reqs := []int32{1, 5, 3, 6, 2, 20}

	waitc := make(chan struct{})

	go func (){
		for _, req := range reqs{
			log.Printf("Sending request: %d\n", req)
			stream.Send(&pb.MaxRequest{
				Res: req,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func (){
		for {
			res, err := stream.Recv()

			if err == io.EOF{
				break
			}
			if err != nil{
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Recived: %v\n", res.Res)
		}
		close(waitc)
	}()

	<- waitc
}