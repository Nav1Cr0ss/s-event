package main

import (
	"github.com/Nav1Cr0ss/s-event/config"
	db "github.com/Nav1Cr0ss/s-event/internal/adapters/repository"
	repository "github.com/Nav1Cr0ss/s-event/internal/adapters/repository/sqlc"
	"github.com/Nav1Cr0ss/s-event/internal/app"
	handler "github.com/Nav1Cr0ss/s-event/internal/ports/grpc"
	pbevent "github.com/Nav1Cr0ss/s-event/pkg/s-design/events_proto/gen/grpc"
	"github.com/Nav1Cr0ss/s-event/server"
	"github.com/Nav1Cr0ss/s-lib/logger"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func main() {
	fx.New(
		fx.Provide(
			config.NewConfiguration,
			func(cfg *config.Config) *logger.Logger {
				return logger.NewLogger(cfg.Debug)
			},
			fx.Annotate(
				db.NewDB,
				fx.As(new(repository.DBTX)),
			),
			fx.Annotate(
				repository.New,
				fx.As(new(repository.Querier)),
			),
			app.NewApplication,
			server.StartListen,

			fx.Annotate(
				grpc.NewServer,
				fx.As(new(grpc.ServiceRegistrar)),
			),
			fx.Annotate(
				handler.NewGRPCHandler,
				fx.As(new(pbevent.EventServiceServer)),
			),
		),
		fx.Invoke(
			pbevent.RegisterEventServiceServer,
			server.StartServer,
		),
	).Run()

}
