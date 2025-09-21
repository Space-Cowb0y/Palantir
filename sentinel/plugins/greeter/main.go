package main


import (
"context"
"flag"
"net"


"google.golang.org/grpc"
"github.com/Space-Cowb0y/Palantir/api"
)


type agentServer struct { api.UnimplementedAgentServiceServer }


func (s *agentServer) Ping(ctx context.Context, req *api.PingRequest) (*api.PingResponse, error) {
return &api.PingResponse{Reply: "Hello, " + req.Message}, nil
}


func (s *agentServer) RunTask(ctx context.Context, req *api.TaskRequest) (*api.TaskResponse, error) {
return &api.TaskResponse{Result: "Executed: " + req.Task, Success: true}, nil
}


func main() {
addr := flag.String("addr", ":50051", "listen address")
flag.Parse()


lis, err := net.Listen("tcp", *addr)
if err != nil { panic(err) }


grpcServer := grpc.NewServer()
api.RegisterAgentServiceServer(grpcServer, &agentServer{})
if err := grpcServer.Serve(lis); err != nil { panic(err) }
}