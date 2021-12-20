package main

import (
	"context"
	"log"
	"net"

	pb "github.com/valbertoenoc/grpc_calculator/calculator/calculatorpb"

	"google.golang.org/grpc"
)

const (
	port = ":50011"
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (s *server) Add(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResult, error) {
	log.Printf("[DEBUG] Received: %v and %v", in.GetOperand1(), in.GetOperand2())
	return &pb.CalculatorResult{Result: in.GetOperand1() + in.GetOperand2()}, nil
}

func (s *server) Subtract(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResult, error) {
	log.Printf("[DEBUG] Received: %v and %v", in.GetOperand1(), in.GetOperand2())
	return &pb.CalculatorResult{Result: in.GetOperand1() - in.GetOperand2()}, nil
}

func (s *server) Multiply(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResult, error) {
	log.Printf("[DEBUG] Received: %v and %v", in.GetOperand1(), in.GetOperand2())
	return &pb.CalculatorResult{Result: in.GetOperand1() * in.GetOperand2()}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResult, error) {
	log.Printf("[DEBUG] Received: %v and %v", in.GetOperand1(), in.GetOperand2())
	return &pb.CalculatorResult{Result: in.GetOperand1() / in.GetOperand2()}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}