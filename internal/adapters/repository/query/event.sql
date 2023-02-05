-- name: CreateEvent :one
INSERT INTO "event" (author, title, description, "type")
VALUES ($1, $2, $3, $4)
RETURNING "id";

-- name: GetEvent :one
SELECT *
FROM "event" as ev
         JOIN event_settings es on ev.id = es.event_id
WHERE ev.id = $1
LIMIT 1;
