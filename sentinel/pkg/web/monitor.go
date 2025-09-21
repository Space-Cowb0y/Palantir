package web

import (
	"context"
	"fmt"
	"net"

	"github.com/Space-Cowb0y/Palantir/sentinel/internal/config"
	"github.com/Space-Cowb0y/Palantir/sentinel/internal/logging"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	cfg *config.Config
	log logging.Logger
	addr string
}

func NewGRPCServer(cfg *config.Config, log logging.Logger) *GRPCServer {
	return &GRPCServer{cfg: cfg, log: log, addr: cfg.GRPC.Listen}
}

func (g *GRPCServer) Addr() string { return g.addr }

func (g *GRPCServer) Start() error {
	lis, err := net.Listen("tcp", g.addr)
	if err != nil { return err }
	grpcServer := grpc.NewServer()
	g.log.Info("gRPC listening", "addr", g.addr)
	return grpcServer.Serve(lis)
}

func (g *GRPCServer) Shutdown(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}