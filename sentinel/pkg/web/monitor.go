package web
import(
  "net"; "google.golang.org/grpc"
  "github.com/Space-Cowb0y/Palantir/sentinel/internal/config"
  "github.com/Space-Cowb0y/Palantir/sentinel/internal/logging"
)

type GRPCServer struct{ cfg *config.Config; log logging.Logger; addr string }
func NewGRPCServer(cfg *config.Config, log logging.Logger)*GRPCServer{ return &GRPCServer{cfg:cfg,log:log,addr:cfg.GRPC.Listen} }
func (g *GRPCServer) Addr() string{ return g.addr }
func (g *GRPCServer) Start() error{ lis,err:=net.Listen("tcp",g.addr); if err!=nil{return err}; srv:=grpc.NewServer(); return srv.Serve(lis) }