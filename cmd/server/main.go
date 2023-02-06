package main

import (
	"github.com/Nav1Cr0ss/s-event/config"
	db "github.com/Nav1Cr0ss/s-event/internal/adapters/repository"
	repository "github.com/Nav1Cr0ss/s-event/internal/adapters/repository/sqlc"
	"github.com/Nav1Cr0ss/s-event/internal/app"
	handler "github.com/Nav1Cr0ss/s-event/internal/ports/grpc"
	"github.com/Nav1Cr0ss/s-event/pkg/s-design/pbevent/gen/pbevent"
	"github.com/Nav1Cr0ss/s-event/server"
	"github.com/Nav1Cr0ss/s-lib/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {

	fx.New(
		fx.Provide(
			config.NewConfiguration,
			func(cfg *config.Config) (*logger.Logger, *zap.Logger) {
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
				func(log *zap.Logger) *grpc.Server {
					options := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
						grpc_ctxtags.UnaryServerInterceptor(),
						grpc_zap.UnaryServerInterceptor(log),
						grpc_validator.UnaryServerInterceptor(),
						grpc_recovery.UnaryServerInterceptor(),
					))
					return grpc.NewServer(options)
				},
				fx.As(new(grpc.ServiceRegistrar)),
			),
			fx.Annotate(
				handler.NewGRPCHandler,
				fx.As(new(pbevent.EventServiceServer)),
			),
		),
		fx.Invoke(
			func(s grpc.ServiceRegistrar, srv pbevent.EventServiceServer, log *logger.Logger) {
				log.Info("start app")
				pbevent.RegisterEventServiceServer(s, srv)
			},
			server.StartServer,
		),
	).Run()

}
