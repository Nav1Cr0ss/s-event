-- name: CreateEvent :one
INSERT INTO "event" (author_id, title, description, "type")
VALUES ($1, $2, $3, $4)
RETURNING "id";

-- name: GetEvent :one
SELECT ev.description,
       ev.title,
       ev.author_id,
       ev.created_at,
       ev.type,
       es.max_participants,
       es.min_participants,
       es.visibility
FROM "event" as ev
         JOIN event_settings es on ev.id = es.event_id
WHERE ev.id = $1
LIMIT 1;
