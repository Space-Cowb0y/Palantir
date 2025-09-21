package plugin

import (
	"errors"
	"net"
	"os/exec"
	"time"

	"github.com/Space-Cowb0y/Palantir/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Loader struct{ conns []*grpc.ClientConn }

func NewLoader() *Loader { return &Loader{} }

func pickFreePort() (string, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	defer l.Close()
	return l.Addr().String(), nil
}

func (l *Loader) StartAgent(binPath string) (api.AgentServiceClient, string, error) {
	addr, err := pickFreePort()
	if err != nil {
		return nil, "", err
	}

	cmd := exec.Command(binPath, "--addr", addr)
	if err := cmd.Start(); err != nil {
		return nil, "", err
	}

	// tentativa de conexão com backoff leve
	var conn *grpc.ClientConn
	for i := 0; i < 20; i++ {
		conn, err = grpc.Dial(addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
			grpc.WithTimeout(300*time.Millisecond),
		)
		if err == nil {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	if conn == nil {
		return nil, "", errors.New("agent not reachable")
	}

	l.conns = append(l.conns, conn)
	zap.L().Info("agent connected", zap.String("addr", addr))
	return api.NewAgentServiceClient(conn), addr, nil
}

func (l *Loader) Cleanup() {
	for _, c := range l.conns {
		_ = c.Close()
	}
}
