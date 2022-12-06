package main

import (
	"io"
	"log"

	pb "github.com/dionofrizal88/go-grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error{
	log.Printf("maxValue was invoked")
	var maximum int32 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF{
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		
		if number := req.Res; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{
				Res: maximum,
			})

			if err != nil{
				log.Fatalf("Error while sending data from client: %v\n", err)
			}
		}
	}
}