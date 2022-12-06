package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/dionofrizal88/go-grpc/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error{
	fmt.Println("Average function was invoked")

	var num float64
	counter := 0
	var average float64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average = num/float64(counter)
			return stream.SendAndClose(&pb.FloatResponse{
				Res: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		num += req.Number
		counter++
	}
}