// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: day_reports.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createDayReport = `-- name: CreateDayReport :one
INSERT INTO day_reports(
    way_uuid,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3
) RETURNING uuid, way_uuid, created_at, updated_at
`

type CreateDayReportParams struct {
	WayUuid   pgtype.UUID      `json:"way_uuid"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) CreateDayReport(ctx context.Context, arg CreateDayReportParams) (DayReport, error) {
	row := q.db.QueryRow(ctx, createDayReport, arg.WayUuid, arg.CreatedAt, arg.UpdatedAt)
	var i DayReport
	err := row.Scan(
		&i.Uuid,
		&i.WayUuid,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDayReportsByRankRange = `-- name: GetDayReportsByRankRange :many
WITH ranked_reports AS (
    SELECT
        dense_rank() OVER (ORDER BY date(day_reports.created_at) DESC) AS rank,
        day_reports.uuid, day_reports.way_uuid, day_reports.created_at, day_reports.updated_at,
        ways.name AS way_name
    FROM day_reports
    INNER JOIN ways ON day_reports.way_uuid = ways.uuid
    WHERE ways.uuid = ANY ($3::UUID[])
),
max_rank_cte AS (
    SELECT
        COALESCE(MAX(rank), 0)::INTEGER AS max_rank
    FROM ranked_reports
)
SELECT ranked_reports.rank, ranked_reports.uuid, ranked_reports.way_uuid, ranked_reports.created_at, ranked_reports.updated_at, ranked_reports.way_name, max_rank_cte.max_rank
FROM ranked_reports, max_rank_cte
WHERE rank BETWEEN ($1::INTEGER) AND ($2::INTEGER)
ORDER BY rank
`

type GetDayReportsByRankRangeParams struct {
	StartRankRange int32         `json:"start_rank_range"`
	EndRankRange   int32         `json:"end_rank_range"`
	WayUuids       []pgtype.UUID `json:"way_uuids"`
}

type GetDayReportsByRankRangeRow struct {
	Rank      int64            `json:"rank"`
	Uuid      pgtype.UUID      `json:"uuid"`
	WayUuid   pgtype.UUID      `json:"way_uuid"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	WayName   string           `json:"way_name"`
	MaxRank   int32            `json:"max_rank"`
}

func (q *Queries) GetDayReportsByRankRange(ctx context.Context, arg GetDayReportsByRankRangeParams) ([]GetDayReportsByRankRangeRow, error) {
	rows, err := q.db.Query(ctx, getDayReportsByRankRange, arg.StartRankRange, arg.EndRankRange, arg.WayUuids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetDayReportsByRankRangeRow{}
	for rows.Next() {
		var i GetDayReportsByRankRangeRow
		if err := rows.Scan(
			&i.Rank,
			&i.Uuid,
			&i.WayUuid,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.WayName,
			&i.MaxRank,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDayReportsCountByWayId = `-- name: GetDayReportsCountByWayId :one
SELECT
    COUNT(*) AS day_reports_count
FROM day_reports
WHERE way_uuid = $1
`

func (q *Queries) GetDayReportsCountByWayId(ctx context.Context, wayUuid pgtype.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, getDayReportsCountByWayId, wayUuid)
	var day_reports_count int64
	err := row.Scan(&day_reports_count)
	return day_reports_count, err
}

const getIsUserHavingPermissionsForDayReport = `-- name: GetIsUserHavingPermissionsForDayReport :one
SELECT
    ways.uuid as way_uuid,
    ways.name as way_name,
    EXISTS (
        SELECT 1
        FROM mentor_users_ways
        WHERE mentor_users_ways.way_uuid = ways.uuid
        AND mentor_users_ways.user_uuid = $1
    ) OR ways.owner_uuid = $1 AS is_permission_given
FROM ways
INNER JOIN day_reports ON ways.uuid = day_reports.way_uuid
WHERE day_reports.uuid = $2
`

type GetIsUserHavingPermissionsForDayReportParams struct {
	UserUuid      pgtype.UUID `json:"user_uuid"`
	DayReportUuid pgtype.UUID `json:"day_report_uuid"`
}

type GetIsUserHavingPermissionsForDayReportRow struct {
	WayUuid           pgtype.UUID `json:"way_uuid"`
	WayName           string      `json:"way_name"`
	IsPermissionGiven pgtype.Bool `json:"is_permission_given"`
}

func (q *Queries) GetIsUserHavingPermissionsForDayReport(ctx context.Context, arg GetIsUserHavingPermissionsForDayReportParams) (GetIsUserHavingPermissionsForDayReportRow, error) {
	row := q.db.QueryRow(ctx, getIsUserHavingPermissionsForDayReport, arg.UserUuid, arg.DayReportUuid)
	var i GetIsUserHavingPermissionsForDayReportRow
	err := row.Scan(&i.WayUuid, &i.WayName, &i.IsPermissionGiven)
	return i, err
}

const getListDayReportsByWayUuid = `-- name: GetListDayReportsByWayUuid :many
SELECT uuid, way_uuid, created_at, updated_at
FROM day_reports
WHERE day_reports.way_uuid = $1
ORDER BY day_reports.created_at DESC
LIMIT $3
OFFSET $2
`

type GetListDayReportsByWayUuidParams struct {
	WayUuid       pgtype.UUID `json:"way_uuid"`
	RequestOffset int32       `json:"request_offset"`
	RequestLimit  int32       `json:"request_limit"`
}

func (q *Queries) GetListDayReportsByWayUuid(ctx context.Context, arg GetListDayReportsByWayUuidParams) ([]DayReport, error) {
	rows, err := q.db.Query(ctx, getListDayReportsByWayUuid, arg.WayUuid, arg.RequestOffset, arg.RequestLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []DayReport{}
	for rows.Next() {
		var i DayReport
		if err := rows.Scan(
			&i.Uuid,
			&i.WayUuid,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}