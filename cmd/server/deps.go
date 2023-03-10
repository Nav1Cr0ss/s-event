package main

import (
	"database/sql"

	"github.com/Nav1Cr0ss/s-event/config"
	"github.com/Nav1Cr0ss/s-event/pkg/s-design/pbevent/gen/pbevent"
	"github.com/Nav1Cr0ss/s-lib/configuration"
	"github.com/Nav1Cr0ss/s-lib/database"
	"github.com/Nav1Cr0ss/s-lib/grpc_server"
	"github.com/Nav1Cr0ss/s-lib/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ProvideConfig() *config.Config {
	cfg := &config.Config{}
	configuration.NewConfiguration(cfg)
	return cfg
}

func ProvideLogger(cfg *config.Config) *logger.Logger {
	return logger.NewLogger(cfg.App.Debug)
}

func ProvideDb(c *config.Config, log *logger.Logger) *sql.DB {
	return database.NewDB(c, log)
}

func ProvideGRPCServer(cfg *config.Config, log *logger.Logger) *grpc_server.GRPCServer {
	return grpc_server.NewGRPCServer(cfg, log)
}

func InvokeRegisterService(cfg *config.Config, s *grpc_server.GRPCServer, h pbevent.EventServiceServer) {
	pbevent.RegisterEventServiceServer(s.Reg, h)

	if cfg.GetDebug() {
		ss := s.Reg.(*grpc.Server)
		reflection.Register(ss)
	}
}
