package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "MalwareProtocol/proto/user.proto"
)

const (
	address = "127.0.0.1:10000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cannel := context.WithTimeout(context.Background(), time.Second*10)
	defer cannel()
	rsp, err := c.Login(ctx, &pb.LoginRequest{Username: "select a current_user", Password: "123"})
	if err != nil {
		log.Fatalf("could not login request: %v", err)
	}
	log.Printf("Username: %s  Code: %d   Msg: %s", rsp.GetUsername(), rsp.Code, rsp.Msg)
}
