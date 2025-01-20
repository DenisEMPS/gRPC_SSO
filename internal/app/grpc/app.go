package grpcapp

import (
	"fmt"
	"grpc/internal/grpc/auth"
	authrpc "grpc/internal/grpc/auth"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	GRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int, auth auth.Auth) *App {
	gRPCServer := grpc.NewServer()

	authrpc.Register(gRPCServer, auth)

	return &App{
		log:        log,
		GRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String(("op"), op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := a.GRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String(("op"), op),
		slog.Int("port", a.port))

	a.GRPCServer.GracefulStop()
}
