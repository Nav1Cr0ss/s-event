package handler

import (
	"log"

	repo "github.com/Nav1Cr0ss/s-event/internal/adapters/repository/sqlc"
	"github.com/Nav1Cr0ss/s-event/pkg/s-design/pbevent/gen/pbevent"
	"github.com/Nav1Cr0ss/s-lib/interceptor"
	"github.com/Nav1Cr0ss/s-lib/strings"
	"google.golang.org/protobuf/types/known/timestamppb"
)

import (
	"context"
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

type User struct {
	Id int `json:"id"`
}

func (h GRPCHandler) GetEvent(ctx context.Context, req *pbevent.GetEventRequest) (*pbevent.GetEventResponse, error) {
	var (
		err   error
		event repo.GetEventRow
	)
	//md, _ := metadata.FromIncomingContext(ctx)
	//user := md.Get("user")[0]
	//
	//us := User{}
	//err = json.Unmarshal([]byte(user), &us)
	//if err != nil {
	//	return nil, err
	//}
	user := ctx.Value("User").(interceptor.User)
	_ = user
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
			MaxParticipants: event.MaxParticipants,
			MinParticipants: event.MinParticipants,
			Visibility:      strings.GetArrayOfStrings(event.Visibility),
		},
	}
	return &resp, nil

}
