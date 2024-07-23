// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: p2p_messages.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMessageInP2PRoom = `-- name: CreateMessageInP2PRoom :one
INSERT INTO p2p_messages (owner_uuid, room_uuid, text)
VALUES ($1, $2, $3)
RETURNING owner_uuid, text
`

type CreateMessageInP2PRoomParams struct {
	OwnerUuid pgtype.UUID `json:"owner_uuid"`
	RoomUuid  pgtype.UUID `json:"room_uuid"`
	Text      string      `json:"text"`
}

type CreateMessageInP2PRoomRow struct {
	OwnerUuid pgtype.UUID `json:"owner_uuid"`
	Text      string      `json:"text"`
}

func (q *Queries) CreateMessageInP2PRoom(ctx context.Context, arg CreateMessageInP2PRoomParams) (CreateMessageInP2PRoomRow, error) {
	row := q.db.QueryRow(ctx, createMessageInP2PRoom, arg.OwnerUuid, arg.RoomUuid, arg.Text)
	var i CreateMessageInP2PRoomRow
	err := row.Scan(&i.OwnerUuid, &i.Text)
	return i, err
}

const getP2PMessagesByRoomUUID = `-- name: GetP2PMessagesByRoomUUID :many
SELECT owner_uuid, text FROM p2p_messages
WHERE room_uuid = $1
`

type GetP2PMessagesByRoomUUIDRow struct {
	OwnerUuid pgtype.UUID `json:"owner_uuid"`
	Text      string      `json:"text"`
}

func (q *Queries) GetP2PMessagesByRoomUUID(ctx context.Context, roomUuid pgtype.UUID) ([]GetP2PMessagesByRoomUUIDRow, error) {
	rows, err := q.db.Query(ctx, getP2PMessagesByRoomUUID, roomUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetP2PMessagesByRoomUUIDRow{}
	for rows.Next() {
		var i GetP2PMessagesByRoomUUIDRow
		if err := rows.Scan(&i.OwnerUuid, &i.Text); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
