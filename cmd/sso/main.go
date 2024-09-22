package main

import (
	"log/slog"
	"os"

	"github.com/FelliniFeed/sso/internal/app"
	"github.com/FelliniFeed/sso/internal/config"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main() {
	//TODO: инициализировать объект конфига
	cfg := config.MustLoad()

	//TODO: инициализировать логгер

	log:= setupLogger(cfg.Env)
	
	log.Info("Start application: ", slog.Any("config", cfg))

	application := app.New(log,cfg.GRPC.Port,cfg.StoragePath,cfg.TokenTTL)

	application.GRPCSrv.MustRun()

	//TODO: инициализировать приложение (app)



	//TODO: запустить gRPC-сервер приложения
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}