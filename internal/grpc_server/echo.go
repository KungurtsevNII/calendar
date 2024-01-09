package grpc_server

import (
	"context"

	pb "github.com/KungurtsevNII/calendar/pkg/pb"
)

func (s *CalendarServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.GetMessage()}, nil
}
