package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"

	pan "github.com/milvum/hummingbird/pkg/panopticon"
	pb "github.com/milvum/hummingbird/proto"
)

var hiberAddr = flag.String("hiberAddr", "localhost:8000", "address for hiberBridge service (ip:port)")
var port = flag.Int("port", 8080, "local port to bind for hummingbird grpc server")

type hummingbirdServer struct {
	savedStatusRequests  []*pb.StatusRequest
	savedMessageRequests []*pb.MessageRequest
	mu                   sync.Mutex
	proxy                pan.HiberBridgeClient
}

func newServer(hiberAddr string) *hummingbirdServer {
	fmt.Println("newServer")
	proxy := pan.NewClient(hiberAddr)
	s := hummingbirdServer{proxy: *proxy}
	return &s
}

func (s *hummingbirdServer) CastStatuses(ctx context.Context, statusRequest *pb.StatusRequest) (*pb.Reply, error) {
	fmt.Println("CastStatuses")
	s.mu.Lock()
	defer s.mu.Unlock()
	s.savedStatusRequests = append(s.savedStatusRequests, statusRequest)
	fmt.Println("statuses", s.savedStatusRequests)
	return &pb.Reply{Success: true}, nil
}

func (s *hummingbirdServer) SendMessage(ctx context.Context, messageRequest *pb.MessageRequest) (*pb.Reply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.savedMessageRequests = append(s.savedMessageRequests, messageRequest)
	s.proxy.Uplink(*messageRequest)
	fmt.Println("messages", s.savedMessageRequests)
	return &pb.Reply{Success: true}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("hello world")
	grpcServer := grpc.NewServer()

	// connect
	s := newServer(*hiberAddr)
	pb.RegisterHummingbirdServer(grpcServer, s)
	grpcServer.Serve(lis)
}
