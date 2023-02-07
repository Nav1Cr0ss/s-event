package main

import (
	"database/sql"

	"github.com/Nav1Cr0ss/s-event/config"
	"github.com/Nav1Cr0ss/s-event/pkg/s-design/pbevent/gen/pbevent"
	"github.com/Nav1Cr0ss/s-lib/configuration"
	"github.com/Nav1Cr0ss/s-lib/database"
	"github.com/Nav1Cr0ss/s-lib/grpc_server"
	"github.com/Nav1Cr0ss/s-lib/logger"
	"go.uber.org/zap"
)

func ProvideConfig() *config.Config {
	cfg := &config.Config{}
	configuration.NewConfiguration(cfg)
	return cfg
}

func ProvideLogger(cfg *config.Config) (*logger.Logger, *zap.Logger) {
	return logger.NewLogger(cfg.App.Debug)
}

func ProvideDb(c *config.Config, log *logger.Logger) *sql.DB {
	return database.NewDB(c, log)
}

func ProvideGRPCServer(cfg *config.Config, log *logger.Logger, zapLogger *zap.Logger) *grpc_server.GRPCServer {
	return grpc_server.NewGRPCServer(cfg, log, zapLogger)
}

func InvokeRegisterService(s *grpc_server.GRPCServer, h pbevent.EventServiceServer) {
	pbevent.RegisterEventServiceServer(s.Reg, h)
}
