package server

import (
	"fmt"
	"net"

	"github.com/Nav1Cr0ss/s-event/config"
	"google.golang.org/grpc"
)

func StartListen(cfg *config.Config) (net.Listener, error) {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		return nil, err
	}
	return lis, nil
}

func StartServer(srv grpc.ServiceRegistrar, lis net.Listener) error {
	server, ok := srv.(*grpc.Server)
	if !ok {
		return fmt.Errorf("error on type casting")
	}

	err := server.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
