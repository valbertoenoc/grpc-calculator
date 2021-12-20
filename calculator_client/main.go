package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/valbertoenoc/grpc_calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50011"
	default_operand1 float64 = 1.0
	default_operand2 float64 = 1.0
)

var (
	result *pb.CalculatorResult
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	var	op1 = default_operand1
	var op2 = default_operand2
	var operation = "none"
	// log.Printf("length arguments: %v", len(os.Args))
	if len(os.Args) == 4 {
		op1, _ = strconv.ParseFloat(os.Args[1], 32)
		op2, _ = strconv.ParseFloat(os.Args[2], 32)
		operation = os.Args[3]
		log.Printf("operand1: %v operand2: %v operation: %v", op1, op2, operation)
	} else {
		log.Fatalf("Invalid Arguments.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	switch operation {
	case "+":
		result, err = c.Add(ctx, &pb.CalculatorRequest{Operand1: op1, Operand2: op2})
	case "-":
		result, err = c.Subtract(ctx, &pb.CalculatorRequest{Operand1: op1, Operand2: op2})
	case "*":
		result, err = c.Multiply(ctx, &pb.CalculatorRequest{Operand1: op1, Operand2: op2})
	case "/":
		result, err = c.Divide(ctx, &pb.CalculatorRequest{Operand1: op1, Operand2: op2})
	}
	
	if err != nil {
		log.Fatalf("Could not calculate: %v", err)
	}
	log.Printf("Result: %v", result.GetResult())
}