package schemas

import "github.com/jackc/pgx/v5/pgtype"

type CreateProposalsRow struct {
	UserUuid pgtype.UUID      `json:"user_uuid"`
	CreateAt pgtype.Timestamp `json:"create_at"`
}

type CreateProposalPayload struct {
	UserUuid *string `json:"user_uuid" extensions:"x-nullable"`
	Proposal *string `json:"proposal" extensions:"x-nullable"`
	CreateAt string  `json:"create_At" extensions:"x-nullable"`
}
