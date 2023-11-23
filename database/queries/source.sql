-- name: CreateModelSource :one
INSERT INTO
	root.model_sources(entity_id, source_base_64)
VALUES
	($1, $2)
RETURNING *;

-- name: DeleteModelSourceByEntityId :exec
DELETE FROM
	root.model_sources
WHERE
	entity_id = $1;

-- name: UpdateModelSourceByEntityId :exec
UPDATE
	root.model_sources
SET
	source_base_64 = $2,
	updated_at = now()
WHERE
	entity_id = $1;

-- name: GetModelSourceByEntityId :one
SELECT
	*
FROM
	root.model_sources
WHERE
	entity_id = $1;