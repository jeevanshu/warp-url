package service

import (
	"log"

	"github.com/jeevanshu/warp-url/pb"
	"golang.org/x/net/context"
)

// Service struct
type Service struct {
}

// Health check
func (svc *Service) Health(ctx context.Context, request *pb.HealthRequest) (*pb.HealthResponse, error) {
	log.Printf("Received message: %s", request.Message)
	return &pb.HealthResponse{
		Response: "URL Shortener is up!"}, nil
}
