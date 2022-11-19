package main

import (
	"fmt"
	"log"

	pb "github.com/dionofrizal88/go-grpc/calculator/proto"
)

func (s *Server) SumFactor(request *pb.CalculatorRequest, stream pb.CalculatorService_SumFactorServer) error{
	log.Printf("SumFactor function was invoked with: %v\n", request)
	N := request.OperanOne
	k := request.OperanTwo
	for N > 1{
		if N % k == 0{
			res := fmt.Sprintf("%d", k)
			stream.Send(&pb.FactorResponse{
				Result: res,
			})
			N = N / k
		} else{
			k = k + 1
		}
	}
	return nil
}




// func (s *Server) GreetManyTimes(request *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error{
// 	log.Printf("GreetManyTimes function was invoked with: %v\n", request)

// 	for i := 0; i < 10; i++{
// 		res := fmt.Sprintf("Hello %s, number %d", request.FirstName, i)
// 		stream.Send(&pb.GreetResponse{
// 			Result: res,
// 		})
// 	}

// 	return nil
// }