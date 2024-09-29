package services

import (
	db "mwcustdev/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	CustdevService *CustdevService
}

func NewService(pool *pgxpool.Pool) *Service {
	queries := db.New(pool)

	return &Service{
		CustdevService: NewCustDevServices(pool, queries),
	}
}
