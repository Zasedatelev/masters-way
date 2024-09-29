package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProposals = `-- name: AddPropodsal :one
INSERT INTO proposals(user_id, proposal, create_at)
VALUES ($1, $2, $3)
RETURNING user_id`

type CreateProposalParams struct {
	UserUuid pgtype.UUID      `json:"user_uuid"`
	Proposal pgtype.Text      `json:"proposals"`
	CreateAt pgtype.Timestamp `json:"create_at"`
}

type CreateProposalRow struct {
	UserUuid pgtype.UUID      `json:"user_uuid"`
	CreateAt pgtype.Timestamp `json:"create_at"`
}

func (q *Queries) CreateProposal(ctx context.Context, arg CreateProposalParams) error {
	row := q.db.QueryRow(ctx, createProposals,
		arg.UserUuid,
		arg.Proposal,
		arg.CreateAt,
	)

	var i CreateProposalRow

	err := row.Scan(&i.UserUuid, &i.CreateAt)
	if err != nil {
		return err
	}

	return nil
}
