package main

import pb "github.com/vikashparashar/gRPC_Golang/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
