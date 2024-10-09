-- name: GetUser :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CountUsersWithEmail :one
SELECT COUNT(*)
FROM users
WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (
  id, username, email, password
) VALUES (
  $1, $2, $3, $4
);
