package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	"sso1/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	log       *slog.Logger
	grpcSerer *grpc.Server
	port      int
}

func New(
	log *slog.Logger,
	port int,
) *App {
	grpcServer := grpc.NewServer()

	auth.Register(grpcServer)
	return &App{
		log:       log,
		grpcSerer: grpcServer,
		port:      port,
	}
}

func (a *App) Run() error {
	const op = "grpc.Run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, er := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if er != nil {
		return fmt.Errorf("%s: %w", op, er)
	}

	log.Info("GRPC Server is running on port", slog.String("adr", l.Addr().String()))

	if er := a.grpcSerer.Serve(l); er != nil {
		return fmt.Errorf("%s: %w", op, er)
	}

	return nil

}
