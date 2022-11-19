package main

import (
	"context"
	"log"

	pb "github.com/dionofrizal88/go-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, request *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Sum function was invoked with %v\n", request)
	result := request.OperanOne + request.OperanTwo
	return &pb.CalculatorResponse{
		Result: result,
	}, nil
}