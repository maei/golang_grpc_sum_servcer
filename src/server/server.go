package main

import (
	"context"
	"github.com/maei/golang_grpc_sum_servcer/src/calculatorpb"
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func main() {
	logger.Info("gRPC sum-calculator-server started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Error("error while listening gRPC Server", err)
		log.Printf("error while listening gRPC Server %v", err)
	}
	s := grpc.NewServer()
	calculatropb.RegisterCalculatorServiceServer(s, &server{})

	errServer := s.Serve(lis)
	if errServer != nil {
		logger.Error("error while serve gRPC Server", errServer)
		log.Printf("error while serve gRPC Server %v", errServer)
	}
}

func (s *server) Calc(ctx context.Context, req *calculatropb.CalculatorRequest) (*calculatropb.CalculatorResponse, error) {
	a := req.GetCalulation().GetA()
	b := req.GetCalulation().GetB()
	result := a + b

	res := &calculatropb.CalculatorResponse{Result: result}
	return res, nil
}