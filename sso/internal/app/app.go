package app

import (
	"log/slog"
	grpcapp "sso1/internal/app/grpc"
	"time"

	_ "google.golang.org/grpc"
)

type App struct {
	grpcSerer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTl time.Duration,
) *App {
	grpcapp := grpcapp.New(log, grpcPort)
	return &App{
		grpcSerer: grpcapp,
	}
}
