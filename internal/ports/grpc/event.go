package handler

import (
	"context"
	"log"

	repo "github.com/Nav1Cr0ss/s-event/internal/adapters/repository/sqlc"
	pbevent "github.com/Nav1Cr0ss/s-event/pkg/s-design/events_proto/gen/grpc"
	"github.com/Nav1Cr0ss/s-lib/enum"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h GRPCHandler) CreateEvent(ctx context.Context, req *pbevent.CreateEventRequest) (*pbevent.CreateEventResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	qs := repo.EventTypeEnum(req.Type)

	event := repo.CreateEventParams{
		AuthorID:    req.AuthorId,
		Title:       req.Title,
		Description: req.Description,
		Type:        qs,
	}
	eventId, err := h.a.CreateEvent(ctx, event)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	resp := pbevent.CreateEventResponse{
		EventId: eventId,
	}
	return &resp, nil

}

func GetArrayOfStrings[T ~string](A []T) []string {
	var tmp []string

	for _, val := range A {
		tmp = append(tmp, string(val))
	}

	return tmp
}
func (h GRPCHandler) GetEvent(ctx context.Context, req *pbevent.GetEventRequest) (*pbevent.GetEventResponse, error) {
	var (
		err   error
		event repo.GetEventRow
	)
	err = req.Validate()
	if err != nil {
		return nil, err
	}

	event, err = h.a.GetEvent(ctx, req.EventId)
	if err != nil {
		return nil, err
	}

	resp := pbevent.GetEventResponse{
		Event: &pbevent.Event{
			//Id:          0,
			AuthorId:    event.AuthorID,
			Title:       event.Title,
			Description: event.Description,
			Type:        string(event.Type),
			CreatedAt:   timestamppb.New(event.CreatedAt),
		},
		EventSettings: &pbevent.EventSettings{
			//Id:              123,
			//EventId:         event.,
			MaxParticipants: event.MaxParticipants,
			MinParticipants: event.MinParticipants,
			Visibility:      GetArrayOfStrings(event.Visibility),
		},
	}
	return &resp, nil

}
