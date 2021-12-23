package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/valbertoenoc/grpc_calculator/proto/calculator"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func NewServer() *server {
	return &server{}
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
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	log.Printf("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	
	gwmux := runtime.NewServeMux()

	err = pb.RegisterCalculatorServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr: ":8090",
		Handler: gwmux,
	}

	log.Println(("Serving gRPC-Gateway on http://0.0.0.0:8090"))
	log.Fatalln(gwServer.ListenAndServe())
}