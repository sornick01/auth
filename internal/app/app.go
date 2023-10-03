package app

import (
	"net"

	"github.com/sornick01/auth/internal/api/user"
	pb "github.com/sornick01/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// App app
type App struct {
	grpcServer *grpc.Server
}

// New app
func New() *App {
	a := &App{}
	a.initGRPCServer()

	return a
}

// RunApp runs app
func (a *App) RunApp(port string) error {
	return a.runGRPCServer(port)
}

func (a *App) initGRPCServer() {
	a.grpcServer = grpc.NewServer()
	reflection.Register(a.grpcServer)
	pb.RegisterUserV1Server(a.grpcServer, user.NewService())
}

func (a *App) runGRPCServer(port string) error {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(l)
	if err != nil {
		return err
	}

	return nil
}
