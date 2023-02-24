package main

import (
	"context"
	"flag"
	"google.golang.org/grpc/encoding/gzip"
	"log"
	"time"

	pb "MalwareProtocol/proto/user.proto"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

var (
	addr = flag.String("addr", "localhost:10000", "the address to connect to")
)

func main() {
	flag.Parse()

	//gzip
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cannel := context.WithTimeout(context.Background(), time.Second*10)
	defer cannel()

	a := pb.Person{
		Name:               "",
		Id:                 0,
		Email:              nil,
		Phones:             nil,
		WeightRecentMonths: nil,
	}

	b := make([]*pb.Person, 10)
	b = append(b, &a)

	rsp, err := c.Login(ctx, &pb.LoginRequest{Username: "select a current_user", Password: "123", People: b})
	if err != nil {
		log.Fatalf("could not login request: %v", err)
	}
	log.Printf("Username: %s  Code: %d   Msg: %s", rsp.GetUsername(), rsp.Code, rsp.Msg)
}
