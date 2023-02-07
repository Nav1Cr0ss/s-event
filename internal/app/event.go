package app

import (
	"context"

	repo "github.com/Nav1Cr0ss/s-event/internal/adapters/repository/sqlc"
)

func (a Application) CreateEvent(ctx context.Context, arg repo.CreateEventParams) (int32, error) {

	return a.repo.CreateEvent(ctx, arg)

}

func (a Application) GetEvent(ctx context.Context, id int32) (repo.GetEventRow, error) {

	return a.repo.GetEvent(ctx, id)
}
