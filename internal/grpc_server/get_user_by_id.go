package grpc_server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/KungurtsevNII/calendar/pkg/pb"
)

func (s *CalendarServer) GetUserByID(
	ctx context.Context,
	req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {

	result, err := s.useCases.GetUserByID(ctx, req.GetUserId())
	switch {
	case err != nil:
		return nil, status.Errorf(codes.Internal, "failed to get user by id: %s", err.Error())
	}

	return &pb.GetUserByIDResponse{
		UserId:     result.ID.String(),
		Email:      "",
		FirstName:  "",
		LastName:   "",
		Patronymic: "",
	}, nil
}
