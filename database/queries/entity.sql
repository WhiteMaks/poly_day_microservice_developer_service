-- name: CreateEntity :one
INSERT INTO
	root.entities(name, type, user_id, parent_entity_id)
VALUES
	($1, $2, $3, $4)
RETURNING *;

-- name: DeleteEntityById :exec
DELETE FROM
	root.entities
WHERE
	id = $1;

-- name: GetEntities :many
SELECT
    *
FROM
    root.entities;