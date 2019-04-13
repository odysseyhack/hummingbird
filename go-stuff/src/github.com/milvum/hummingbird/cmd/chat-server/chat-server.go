package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/milvum/hummingbird/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

type hummingbirdServer struct {
	savedStatusRequests  []*pb.StatusRequest
	savedMessageRequests []*pb.MessageRequest
	mu                   sync.Mutex
}

func newServer() *hummingbirdServer {
	s := &hummingbirdServer{}
	return s
}

func (s *hummingbirdServer) CastStatuses(ctx context.Context, statusRequest *pb.StatusRequest) (*pb.Reply, error) {
	return &pb.Reply{Success: true}, nil
}

func (s *hummingbirdServer) SendMessage(ctx context.Context, messageRequest *pb.MessageRequest) (*pb.Reply, error) {
	return &pb.Reply{Success: true}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("hello world")
	grpcServer := grpc.NewServer()
	// connect
	pb.RegisterHummingbirdServer(grpcServer, newServer())

	grpcServer.Serve(lis)
}
