package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"syscall"

	"github.com/ds248a/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/KungurtsevNII/calendar/internal/config"
	"github.com/KungurtsevNII/calendar/internal/grpc_server"
	"github.com/KungurtsevNII/calendar/internal/infrastructure/mongodb"
	"github.com/KungurtsevNII/calendar/internal/use_cases"
	pb "github.com/KungurtsevNII/calendar/pkg/pb"
)

func main() {
	ctx := context.Background()

	cfg := config.MustLoad()

	mongoCli := initMongo(ctx, cfg.Mongo)

	useCases, err := use_cases.New(mongoCli)
	if err != nil {
		log.Fatalf("can't create use cases service: %s", err.Error())
	}

	initAndStartGRPCServer(cfg.GRPC, useCases)
}

func initMongo(ctx context.Context, cfg config.MongoConfig) *mongodb.MongoDB {
	db, err := mongodb.New(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create mongo db driver: %s", err.Error())
	}

	closer.Add(func() {
		_ = db.Disconnect(ctx)
	})

	return db
}

func initAndStartGRPCServer(cfg config.GRPCServerConfig, service *use_cases.Service) {
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %s", err.Error())
	}

	calendarImpl := grpc_server.New(service)
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterCalendarServer(grpcServer, calendarImpl)
	log.Printf("gRPC server listening at %s", grpcListener.Addr())

	closer.Add(grpcServer.GracefulStop)
	go closer.ListenSignal(syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	if err := grpcServer.Serve(grpcListener); err != nil {
		log.Fatalf("failed to start grpc server: %s", err.Error())
	}
}
