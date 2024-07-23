// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddUserToGroupRoom(ctx context.Context, arg AddUserToGroupRoomParams) error
	AddUserToP2PRoom(ctx context.Context, arg AddUserToP2PRoomParams) error
	CreateGroupRoom(ctx context.Context, arg CreateGroupRoomParams) (CreateGroupRoomRow, error)
	CreateMessageInGroupRoom(ctx context.Context, arg CreateMessageInGroupRoomParams) (CreateMessageInGroupRoomRow, error)
	CreateMessageInP2PRoom(ctx context.Context, arg CreateMessageInP2PRoomParams) (CreateMessageInP2PRoomRow, error)
	CreateP2PRoom(ctx context.Context, createdAt pgtype.Timestamp) (CreateP2PRoomRow, error)
	GetGroupMessagesByRoomUUID(ctx context.Context, roomUuid pgtype.UUID) ([]GetGroupMessagesByRoomUUIDRow, error)
	GetGroupRoomByUUID(ctx context.Context, arg GetGroupRoomByUUIDParams) (GetGroupRoomByUUIDRow, error)
	GetGroupRoomsByUserUUID(ctx context.Context, userUuid pgtype.UUID) ([]GetGroupRoomsByUserUUIDRow, error)
	GetP2PMessagesByRoomUUID(ctx context.Context, roomUuid pgtype.UUID) ([]GetP2PMessagesByRoomUUIDRow, error)
	GetP2PRoomByUUID(ctx context.Context, arg GetP2PRoomByUUIDParams) (GetP2PRoomByUUIDRow, error)
	GetP2PRoomsWithInterlocutorByUserUUID(ctx context.Context, userUuid pgtype.UUID) ([]GetP2PRoomsWithInterlocutorByUserUUIDRow, error)
	ToggleBlockGroupRoom(ctx context.Context, arg ToggleBlockGroupRoomParams) error
	ToggleBlockP2PRoom(ctx context.Context, arg ToggleBlockP2PRoomParams) error
}

var _ Querier = (*Queries)(nil)
