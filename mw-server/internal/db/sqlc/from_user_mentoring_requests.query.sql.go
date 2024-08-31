// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: from_user_mentoring_requests.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFromUserMentoringRequest = `-- name: CreateFromUserMentoringRequest :one
INSERT INTO from_user_mentoring_requests(
    user_uuid,
    way_uuid
) VALUES (
    $1, $2
) RETURNING user_uuid, way_uuid
`

type CreateFromUserMentoringRequestParams struct {
	UserUuid pgtype.UUID `json:"user_uuid"`
	WayUuid  pgtype.UUID `json:"way_uuid"`
}

func (q *Queries) CreateFromUserMentoringRequest(ctx context.Context, arg CreateFromUserMentoringRequestParams) (FromUserMentoringRequest, error) {
	row := q.db.QueryRow(ctx, createFromUserMentoringRequest, arg.UserUuid, arg.WayUuid)
	var i FromUserMentoringRequest
	err := row.Scan(&i.UserUuid, &i.WayUuid)
	return i, err
}

const deleteFromUserMentoringRequest = `-- name: DeleteFromUserMentoringRequest :exec
DELETE FROM from_user_mentoring_requests
WHERE user_uuid = $1 AND way_uuid = $2
`

type DeleteFromUserMentoringRequestParams struct {
	UserUuid pgtype.UUID `json:"user_uuid"`
	WayUuid  pgtype.UUID `json:"way_uuid"`
}

func (q *Queries) DeleteFromUserMentoringRequest(ctx context.Context, arg DeleteFromUserMentoringRequestParams) error {
	_, err := q.db.Exec(ctx, deleteFromUserMentoringRequest, arg.UserUuid, arg.WayUuid)
	return err
}

const getFromUserMentoringRequestWaysByUserId = `-- name: GetFromUserMentoringRequestWaysByUserId :many
SELECT
    ways.uuid,
    ways.name,
    ways.goal_description,
    ways.updated_at,
    ways.created_at,
    ways.estimation_time,
    ways.owner_uuid,
    ways.copied_from_way_uuid,
    ways.is_completed,
    ways.is_private,
    (SELECT COUNT(*) FROM metrics WHERE metrics.way_uuid = ways.uuid) AS way_metrics_total,
    (SELECT COUNT(*) FROM metrics WHERE metrics.way_uuid = ways.uuid AND metrics.is_done = true) AS way_metrics_done,
    (SELECT COUNT(*) FROM favorite_users_ways WHERE favorite_users_ways.way_uuid = ways.uuid) AS way_favorite_for_users,
    (SELECT COUNT(*) FROM day_reports WHERE day_reports.way_uuid = ways.uuid) AS way_day_reports_amount,
    COALESCE(
        ARRAY(
            SELECT composite_ways.child_uuid
            FROM composite_ways
            WHERE composite_ways.parent_uuid = ways.uuid
        ),
        '{}'
    )::VARCHAR[] AS children_uuids
FROM from_user_mentoring_requests
JOIN ways
    ON $1 = from_user_mentoring_requests.user_uuid
    AND from_user_mentoring_requests.way_uuid = ways.uuid
`

type GetFromUserMentoringRequestWaysByUserIdRow struct {
	Uuid                pgtype.UUID      `json:"uuid"`
	Name                string           `json:"name"`
	GoalDescription     string           `json:"goal_description"`
	UpdatedAt           pgtype.Timestamp `json:"updated_at"`
	CreatedAt           pgtype.Timestamp `json:"created_at"`
	EstimationTime      int32            `json:"estimation_time"`
	OwnerUuid           pgtype.UUID      `json:"owner_uuid"`
	CopiedFromWayUuid   pgtype.UUID      `json:"copied_from_way_uuid"`
	IsCompleted         bool             `json:"is_completed"`
	IsPrivate           bool             `json:"is_private"`
	WayMetricsTotal     int64            `json:"way_metrics_total"`
	WayMetricsDone      int64            `json:"way_metrics_done"`
	WayFavoriteForUsers int64            `json:"way_favorite_for_users"`
	WayDayReportsAmount int64            `json:"way_day_reports_amount"`
	ChildrenUuids       []string         `json:"children_uuids"`
}

func (q *Queries) GetFromUserMentoringRequestWaysByUserId(ctx context.Context, userUuid pgtype.UUID) ([]GetFromUserMentoringRequestWaysByUserIdRow, error) {
	rows, err := q.db.Query(ctx, getFromUserMentoringRequestWaysByUserId, userUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFromUserMentoringRequestWaysByUserIdRow{}
	for rows.Next() {
		var i GetFromUserMentoringRequestWaysByUserIdRow
		if err := rows.Scan(
			&i.Uuid,
			&i.Name,
			&i.GoalDescription,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.EstimationTime,
			&i.OwnerUuid,
			&i.CopiedFromWayUuid,
			&i.IsCompleted,
			&i.IsPrivate,
			&i.WayMetricsTotal,
			&i.WayMetricsDone,
			&i.WayFavoriteForUsers,
			&i.WayDayReportsAmount,
			&i.ChildrenUuids,
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

const getFromUserMentoringRequestWaysByWayId = `-- name: GetFromUserMentoringRequestWaysByWayId :many
SELECT
    users.uuid, users.name, users.email, users.description, users.created_at, users.image_url, users.is_mentor
FROM from_user_mentoring_requests
JOIN users ON from_user_mentoring_requests.user_uuid = users.uuid
WHERE from_user_mentoring_requests.way_uuid = $1
`

func (q *Queries) GetFromUserMentoringRequestWaysByWayId(ctx context.Context, wayUuid pgtype.UUID) ([]User, error) {
	rows, err := q.db.Query(ctx, getFromUserMentoringRequestWaysByWayId, wayUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Uuid,
			&i.Name,
			&i.Email,
			&i.Description,
			&i.CreatedAt,
			&i.ImageUrl,
			&i.IsMentor,
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