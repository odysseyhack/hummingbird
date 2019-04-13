package chat

import (
	"context"
	"sync"

	pb "github.com/milvum/hummingbird/proto"
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
	s.mu.Lock()
	defer s.mu.Unlock()
	s.savedStatusRequests = append(s.savedStatusRequests, statusRequest)
	return &pb.Reply{Success: true}, nil
}

func (s *hummingbirdServer) SendMessage(ctx context.Context, messageRequest *pb.MessageRequest) (*pb.Reply, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.savedMessageRequests = append(s.savedMessageRequests, messageRequest)
	return &pb.Reply{Success: true}, nil
}
