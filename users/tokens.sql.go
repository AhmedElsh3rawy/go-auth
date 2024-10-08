// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tokens.sql

package users

import (
	"context"

	"github.com/google/uuid"
)

const addToken = `-- name: AddToken :exec
INSERT INTO tokens (user_id, token)
VALUES ($1, $2)
`

type AddTokenParams struct {
	UserID uuid.NullUUID
	Token  string
}

func (q *Queries) AddToken(ctx context.Context, arg AddTokenParams) error {
	_, err := q.db.ExecContext(ctx, addToken, arg.UserID, arg.Token)
	return err
}

const deleteToken = `-- name: DeleteToken :exec
DELETE FROM tokens
WHERE user_id = $1
`

func (q *Queries) DeleteToken(ctx context.Context, userID uuid.NullUUID) error {
	_, err := q.db.ExecContext(ctx, deleteToken, userID)
	return err
}
