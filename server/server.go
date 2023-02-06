package server

import (
	"fmt"
	"net"

	"github.com/Nav1Cr0ss/s-event/config"
	"github.com/Nav1Cr0ss/s-lib/logger"
	"google.golang.org/grpc"
)

func StartListen(cfg *config.Config, log *logger.Logger) net.Listener {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		_ = lis.Close()
		log.Fatalf("error on starting listening : %s", err)
	}
	return lis
}

func StartServer(cfg *config.Config, srv grpc.ServiceRegistrar, lis net.Listener, log *logger.Logger) {
	server, ok := srv.(*grpc.Server)
	if !ok {
		log.Fatal("error on starting serving")
	}

	log.Infof("Listening tcp: %s", lis.Addr())
	log.Infof("Debug: %t", cfg.Debug)

	err := server.Serve(lis)
	if err != nil {
		_ = lis.Close()
		log.Fatal("error on starting serving")
	}

}
