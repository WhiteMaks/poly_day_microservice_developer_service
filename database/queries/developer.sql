-- name: CreateDeveloper :one
INSERT INTO
    root.developers(user_id)
VALUES
    ($1)
RETURNING *;

-- name: GetMaxUserIdFromDevelopers :one
SELECT
    *
FROM
    root.developers
WHERE
    user_id = (
        SELECT
            MAX(user_id)
        FROM
            root.developers
    );