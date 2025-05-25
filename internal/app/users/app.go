package usersapp

import (
	"fmt"
	"net"

	userserver "github.com/wrtgvr/grpc-microblog/internal/grpc/users"
	userservice "github.com/wrtgvr/grpc-microblog/internal/services/users"
	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func NewApp(port int) *App {
	gRPCServer := grpc.NewServer()
	service := userservice.NewDefaulService()

	userserver.Register(gRPCServer, service)

	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic(err)
	}

	if err := a.gRPCServer.Serve(ln); err != nil {
		return err
	}

	return nil
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
