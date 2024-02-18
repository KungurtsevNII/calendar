package grpc_server

import (
	"github.com/KungurtsevNII/calendar/internal/use_cases"
	pb "github.com/KungurtsevNII/calendar/pkg/pb"
)

type CalendarServer struct {
	useCases *use_cases.Service
	pb.UnimplementedCalendarServer
}

func New(service *use_cases.Service) *CalendarServer {
	return &CalendarServer{
		useCases: service,
	}
}
