package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/KungurtsevNII/calendar/internal/config"
	"github.com/KungurtsevNII/calendar/internal/grpc_server"
	"github.com/KungurtsevNII/calendar/internal/infrastructure/mongodb"
	pb "github.com/KungurtsevNII/calendar/pkg/pb"
)

// protoc calendar.proto --go_out=pkg/pb --go_opt=paths=source_relative --proto_path=api --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative
func main() {
	ctx := context.Background()

	cfg := config.MustLoad()

	mongoCli := initMongo(ctx, cfg.Mongo)
	_ = mongoCli

	initAndStartGRPCServer(cfg.GRPC)
}

func initMongo(ctx context.Context, cfg config.MongoConfig) *mongodb.MongoDB {
	db, err := mongodb.New(ctx, cfg)
	if err != nil {
		log.Fatalf("can't create mongo db driver: %s", err.Error())
	}

	return db
}

func initAndStartGRPCServer(cfg config.GRPCServerConfig) {
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	calendarImpl := grpc_server.New()
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterCalendarServer(grpcServer, calendarImpl)
	log.Printf("gRPC server listening at %v", grpcListener.Addr())

	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
