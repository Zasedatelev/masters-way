// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: favorite_users_ways.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getFavoriteForUserUuidsByWayId = `-- name: GetFavoriteForUserUuidsByWayId :one
SELECT COUNT(*)
FROM favorite_users_ways
WHERE favorite_users_ways.way_uuid = $1
`

func (q *Queries) GetFavoriteForUserUuidsByWayId(ctx context.Context, wayUuid pgtype.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, getFavoriteForUserUuidsByWayId, wayUuid)
	var count int64
	err := row.Scan(&count)
	return count, err
}