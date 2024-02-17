package grpc_server

import (
	"context"

	pb "github.com/KungurtsevNII/calendar/pkg/pb"
)

func (s *CalendarServer) GetUserByID(
	ctx context.Context,
	req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {

	return &pb.GetUserByIDResponse{}, nil
}
