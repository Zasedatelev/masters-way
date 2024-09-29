package services

import (
	"context"
	"errors"
	db "mwcustdev/internal/db/sqlc"
	"time"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrPrivateRoomAlreadyExists = errors.New("A private room for these users already exists!")

type CustdevRepository interface {
	CreateProposal(ctx context.Context, arg db.CreateProposalParams) error
	WithTx(tx pgx.Tx) *db.Queries
}

type CustdevService struct {
	pool              *pgxpool.Pool
	custdevRepository CustdevRepository
}

func NewCustDevServices(pool *pgxpool.Pool, custdevRepository CustdevRepository) *CustdevService {
	return &CustdevService{pool: pool, custdevRepository: custdevRepository}
}

type CreatreProposalsParams struct {
	UserUuid uuid.UUID
	Proposal string
	CreateAt string
}

func (custdevService *CustdevService) CreateProposal(ctx context.Context, param *CreatreProposalsParams) error {
	now := time.Now()
	tx, err := custdevService.pool.Begin(ctx)

	createProposalParam := db.CreateProposalParams{
		UserUuid: pgtype.UUID{Bytes: param.UserUuid, Valid: true},
		Proposal: pgtype.Text{String: param.Proposal, Valid: param.Proposal != ""},
		CreateAt: pgtype.Timestamp{Time: now, Valid: true},
	}

	qtx := custdevService.custdevRepository.WithTx(tx)

	err = qtx.CreateProposal(ctx, createProposalParam)
	if err != nil {
		return err
	}

	return nil
}
