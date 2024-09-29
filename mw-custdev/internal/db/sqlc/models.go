package db

import "github.com/jackc/pgx/v5/pgtype"

type Proposals struct {
	Userid    pgtype.UUID      `json:"user_id"`
	Proposals string           `json:"proposals"`
	CreateAt  pgtype.Timestamp `json:"create_at"`
}
