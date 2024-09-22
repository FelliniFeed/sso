package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/FelliniFeed/sso/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

// Оснвной конструктор
func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	//TODO: инициализировать хранилище (storage)

	//TODO: init auth service( auth)

	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
