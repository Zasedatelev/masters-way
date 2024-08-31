// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: problems.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProblem = `-- name: CreateProblem :one
WITH way_info AS (
    SELECT
        ways.uuid AS way_uuid,
        ways.name AS way_name
    FROM day_reports
    INNER JOIN ways ON ways.uuid = day_reports.way_uuid
    WHERE day_reports.uuid = $6
),
owner_info AS (
    SELECT name AS owner_name
    FROM users
    WHERE uuid = $5
)
INSERT INTO problems(
    created_at,
    updated_at,
    description,
    is_done,
    owner_uuid,
    day_report_uuid
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING uuid, created_at, updated_at, description, is_done, owner_uuid, day_report_uuid,
    (SELECT way_uuid FROM way_info) AS way_uuid,
    (SELECT way_name FROM way_info) AS way_name,
    (SELECT owner_name FROM owner_info) AS owner_name
`

type CreateProblemParams struct {
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	Description   string           `json:"description"`
	IsDone        bool             `json:"is_done"`
	OwnerUuid     pgtype.UUID      `json:"owner_uuid"`
	DayReportUuid pgtype.UUID      `json:"day_report_uuid"`
}

type CreateProblemRow struct {
	Uuid          pgtype.UUID      `json:"uuid"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	Description   string           `json:"description"`
	IsDone        bool             `json:"is_done"`
	OwnerUuid     pgtype.UUID      `json:"owner_uuid"`
	DayReportUuid pgtype.UUID      `json:"day_report_uuid"`
	WayUuid       pgtype.UUID      `json:"way_uuid"`
	WayName       string           `json:"way_name"`
	OwnerName     string           `json:"owner_name"`
}

func (q *Queries) CreateProblem(ctx context.Context, arg CreateProblemParams) (CreateProblemRow, error) {
	row := q.db.QueryRow(ctx, createProblem,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Description,
		arg.IsDone,
		arg.OwnerUuid,
		arg.DayReportUuid,
	)
	var i CreateProblemRow
	err := row.Scan(
		&i.Uuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.IsDone,
		&i.OwnerUuid,
		&i.DayReportUuid,
		&i.WayUuid,
		&i.WayName,
		&i.OwnerName,
	)
	return i, err
}

const deleteProblem = `-- name: DeleteProblem :exec
DELETE FROM problems
WHERE uuid = $1
`

func (q *Queries) DeleteProblem(ctx context.Context, problemUuid pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteProblem, problemUuid)
	return err
}

const getIsUserHavingPermissionsForProblem = `-- name: GetIsUserHavingPermissionsForProblem :one
SELECT
    ways.uuid as way_uuid,
    EXISTS (
        SELECT 1
        FROM mentor_users_ways
        WHERE mentor_users_ways.way_uuid = ways.uuid
        AND mentor_users_ways.user_uuid = $1
    ) OR ways.owner_uuid = $1 AS is_permission_given
FROM ways
INNER JOIN day_reports ON ways.uuid = day_reports.way_uuid
INNER JOIN problems ON problems.day_report_uuid = day_reports.uuid
WHERE problems.uuid = $2
`

type GetIsUserHavingPermissionsForProblemParams struct {
	UserUuid    pgtype.UUID `json:"user_uuid"`
	ProblemUuid pgtype.UUID `json:"problem_uuid"`
}

type GetIsUserHavingPermissionsForProblemRow struct {
	WayUuid           pgtype.UUID `json:"way_uuid"`
	IsPermissionGiven pgtype.Bool `json:"is_permission_given"`
}

func (q *Queries) GetIsUserHavingPermissionsForProblem(ctx context.Context, arg GetIsUserHavingPermissionsForProblemParams) (GetIsUserHavingPermissionsForProblemRow, error) {
	row := q.db.QueryRow(ctx, getIsUserHavingPermissionsForProblem, arg.UserUuid, arg.ProblemUuid)
	var i GetIsUserHavingPermissionsForProblemRow
	err := row.Scan(&i.WayUuid, &i.IsPermissionGiven)
	return i, err
}

const getListProblemsByDayReportId = `-- name: GetListProblemsByDayReportId :many
SELECT uuid, created_at, updated_at, description, is_done, owner_uuid, day_report_uuid FROM problems
WHERE problems.day_report_uuid = $1
ORDER BY created_at
`

func (q *Queries) GetListProblemsByDayReportId(ctx context.Context, dayReportUuid pgtype.UUID) ([]Problem, error) {
	rows, err := q.db.Query(ctx, getListProblemsByDayReportId, dayReportUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Problem{}
	for rows.Next() {
		var i Problem
		if err := rows.Scan(
			&i.Uuid,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Description,
			&i.IsDone,
			&i.OwnerUuid,
			&i.DayReportUuid,
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

const getProblemsByDayReportUuids = `-- name: GetProblemsByDayReportUuids :many
SELECT uuid, created_at, updated_at, description, is_done, owner_uuid, day_report_uuid
FROM problems
WHERE problems.day_report_uuid = ANY($1::UUID[])
`

func (q *Queries) GetProblemsByDayReportUuids(ctx context.Context, dollar_1 []pgtype.UUID) ([]Problem, error) {
	rows, err := q.db.Query(ctx, getProblemsByDayReportUuids, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Problem{}
	for rows.Next() {
		var i Problem
		if err := rows.Scan(
			&i.Uuid,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Description,
			&i.IsDone,
			&i.OwnerUuid,
			&i.DayReportUuid,
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

const updateProblem = `-- name: UpdateProblem :one
WITH way_info AS (
    SELECT
        ways.uuid AS way_uuid,
        ways.name AS way_name
    FROM day_reports
    INNER JOIN ways ON ways.uuid = day_reports.way_uuid
    WHERE day_reports.uuid = (SELECT day_report_uuid FROM problems WHERE uuid = $4)
),
owner_info AS (
    SELECT name AS owner_name
    FROM users
    WHERE uuid = (SELECT owner_uuid FROM problems WHERE uuid = $4)
)
UPDATE problems
SET
updated_at = coalesce($1, updated_at),
description = coalesce($2, description),
is_done = coalesce($3, is_done)
WHERE problems.uuid = $4
RETURNING uuid, created_at, updated_at, description, is_done, owner_uuid, day_report_uuid,
    (SELECT way_uuid FROM way_info) AS way_uuid,
    (SELECT way_name FROM way_info) AS way_name,
    (SELECT owner_name FROM owner_info) AS owner_name
`

type UpdateProblemParams struct {
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	Description pgtype.Text      `json:"description"`
	IsDone      pgtype.Bool      `json:"is_done"`
	ProblemUuid pgtype.UUID      `json:"problem_uuid"`
}

type UpdateProblemRow struct {
	Uuid          pgtype.UUID      `json:"uuid"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	Description   string           `json:"description"`
	IsDone        bool             `json:"is_done"`
	OwnerUuid     pgtype.UUID      `json:"owner_uuid"`
	DayReportUuid pgtype.UUID      `json:"day_report_uuid"`
	WayUuid       pgtype.UUID      `json:"way_uuid"`
	WayName       string           `json:"way_name"`
	OwnerName     string           `json:"owner_name"`
}

func (q *Queries) UpdateProblem(ctx context.Context, arg UpdateProblemParams) (UpdateProblemRow, error) {
	row := q.db.QueryRow(ctx, updateProblem,
		arg.UpdatedAt,
		arg.Description,
		arg.IsDone,
		arg.ProblemUuid,
	)
	var i UpdateProblemRow
	err := row.Scan(
		&i.Uuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.IsDone,
		&i.OwnerUuid,
		&i.DayReportUuid,
		&i.WayUuid,
		&i.WayName,
		&i.OwnerName,
	)
	return i, err
}