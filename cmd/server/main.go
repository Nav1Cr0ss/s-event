package main

import (
	repository "github.com/Nav1Cr0ss/s-event/internal/adapters/repository/sqlc"
	"github.com/Nav1Cr0ss/s-event/internal/app"
	handler "github.com/Nav1Cr0ss/s-event/internal/ports/grpc"
	"github.com/Nav1Cr0ss/s-event/pkg/s-design/pbevent/gen/pbevent"
	"github.com/Nav1Cr0ss/s-lib/grpc_server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			ProvideConfig,
			ProvideLogger,
			fx.Annotate(
				ProvideDb,
				fx.As(new(repository.DBTX)),
			),
			fx.Annotate(
				repository.New,
				fx.As(new(repository.Querier)),
			),
			app.NewApplication,

			fx.Annotate(
				handler.NewGRPCHandler,
				fx.As(new(pbevent.EventServiceServer)),
			),
			ProvideGRPCServer,
		),
		fx.Invoke(
			InvokeRegisterService,
			grpc_server.Serve,
		),
	).Run()
}
