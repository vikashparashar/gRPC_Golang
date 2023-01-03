package main

import pb "github.com/vikashparashar/gRPC_Golang/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
