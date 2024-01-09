package grpc_server

import (
	pb "github.com/KungurtsevNII/calendar/pkg/pb"
)

type CalendarServer struct {
	pb.UnimplementedCalendarServer
}

func New() *CalendarServer {
	return &CalendarServer{}
}
