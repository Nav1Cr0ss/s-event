package handler

import (
	"context"
	"fmt"

	"github.com/Nav1Cr0ss/s-event/internal/app"
	"github.com/Nav1Cr0ss/s-event/pkg/s-design/pbevent/gen/pbevent"
	"github.com/Nav1Cr0ss/s-lib/domains/user"
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

func (h GRPCHandler) getUser(ctx context.Context) (user.User, error) {
	u, ok := ctx.Value("user").(user.User)
	if !ok {
		return u, fmt.Errorf("user doesent parsed")
	}
	return u, nil

}
