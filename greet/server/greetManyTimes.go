package main

import (
	"fmt"
	"log"

	pb "github.com/dionofrizal88/go-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(request *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error{
	log.Printf("GreetManyTimes function was invoked with: %v\n", request)

	for i := 0; i < 10; i++{
		res := fmt.Sprintf("Hello %s, number %d", request.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
