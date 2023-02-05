// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: event.sql

package repository

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO "event" (author, title, description, "type")
VALUES ($1, $2, $3, $4)
RETURNING "id"
`

type CreateEventParams struct {
	Author      string
	Title       string
	Description string
	Type        EventTypeEnum
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createEvent,
		arg.Author,
		arg.Title,
		arg.Description,
		arg.Type,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getEvent = `-- name: GetEvent :one
SELECT ev.id, author, title, description, type, created_at, es.id, event_id, max_participants, min_participants, visibility
FROM "event" as ev
         JOIN event_settings es on ev.id = es.event_id
WHERE ev.id = $1
LIMIT 1
`

type GetEventRow struct {
	ID              int32
	Author          string
	Title           string
	Description     string
	Type            EventTypeEnum
	CreatedAt       time.Time
	ID_2            int32
	EventID         int32
	MaxParticipants int32
	MinParticipants int32
	Visibility      []EventVisibilityEnum
}

func (q *Queries) GetEvent(ctx context.Context, id int32) (GetEventRow, error) {
	row := q.db.QueryRowContext(ctx, getEvent, id)
	var i GetEventRow
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Title,
		&i.Description,
		&i.Type,
		&i.CreatedAt,
		&i.ID_2,
		&i.EventID,
		&i.MaxParticipants,
		&i.MinParticipants,
		pq.Array(&i.Visibility),
	)
	return i, err
}
