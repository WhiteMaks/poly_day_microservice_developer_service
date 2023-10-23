-- name: CreateRelease :one
INSERT INTO
	root.releases(name, user_id)
VALUES
	($1, $2)
RETURNING *;

-- name: UpdateReleaseById :exec
UPDATE
	root.releases
SET
	name = $2,
	updated_at = now()
WHERE
	id = $1;

-- name: DeleteReleaseById :exec
DELETE FROM
	root.releases
WHERE
	id = $1;

-- name: GetReleases :many
SELECT
    *
FROM
    root.releases;