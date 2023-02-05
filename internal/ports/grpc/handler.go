package handler

import (
	"github.com/Nav1Cr0ss/s-event/internal/app"
	pbevent "github.com/Nav1Cr0ss/s-event/pkg/s-design/events_proto/gen/grpc"
	"github.com/Nav1Cr0ss/s-lib/logger"
)

type GRPCHandler struct {
	pbevent.EventServiceServer
	log *logger.Logger
	a   app.Application
}

func NewGRPCHandler(log *logger.Logger, a app.Application) GRPCHandler {
	h := GRPCHandler{
		log: log,
		a:   a,
	}

	return h
}
