package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "MalwareProtocol/proto/user.proto"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received username: %s", in.Username)
	log.Printf("Received password: %s", in.Password)
	return &pb.LoginResponse{Username: in.Username, Age: 18, Code: 0, Msg: "登陆成功"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}
	grpc := grpc.NewServer()
	pb.RegisterUserServiceServer(grpc, &Server{})
	if err := grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
