package main

import (
	"context"
	"log"
	"net"

	pb "github.com/KungurtsevNII/calendar/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CalendarServer struct {
	pb.UnimplementedCalendarServer
}

func (s *CalendarServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.GetMessage()}, nil
}

// protoc calendar.proto --go_out=pkg/pb --go_opt=paths=source_relative --proto_path=api --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative
func main() {
	grpcListener, err := net.Listen("tcp", "127.0.0.1:5001")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	clendarServer := &CalendarServer{}
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterCalendarServer(grpcServer, clendarServer)
	log.Printf("gRPC server listening at %v", grpcListener.Addr())


	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
