-- name: AddToken :exec
INSERT INTO tokens (user_id, token)
VALUES ($1, $2);

-- name: DeleteToken :exec
DELETE FROM tokens
WHERE user_id = $1;
